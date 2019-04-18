package url

func hexToByte(char byte) byte {

    if 'a' <= char && char <= 'f' { return char - 'a' + 10 }
    if 'A' <= char && char <= 'F' { return char - 'A' + 10 }
    if '0' <= char && char <= '9' { return char - '0' }

	return 0
}

func (d *decoder) Decode(data []byte) (match bool, decoded []byte) {

	%% machine decode_url;
	%% write data;

	n := 0
	foundPercent := false
	foundPlus := false
	for i := 0; i < len(data); {

		switch data[i] {
		case '%':
			n++
			foundPercent = true

			if i+2 >= len(data) {
				return false, nil
			}

			for _, d := range []int{1,2} {
				if ! ( '0' <= data[i+d] && data[i+d] <= '9' ||
					'a' <= data[i+d] && data[i+d] <= 'f' ||
					'A' <= data[i+d] && data[i+d] <= 'F') {
					return false, nil
				}
			}

            i += 3
        case '+':
            foundPlus = true
            i++
		default:
			i++
		}
	}

	if !(foundPercent || foundPlus) {
	    return false, nil
	}

	cs, p, pe, eof := 0, 0, len(data), len(data)
	mark := 0
	decoded = make([]byte, len(data)-2*n)
	j := 0
	_ = eof

	%%{
	    action mark { mark = p }
	    action decode { decoded[j] = hexToByte(data[mark:p][0])<<4 | hexToByte(data[mark:p][1]); j++ }
	    action space { decoded[j] = ' '; j++ }
	    action store { decoded[j] = data[p]; j++ }

	    percentEncode = "%" ( xdigit xdigit ) >mark %decode;
	    spaceChar = "+" $space;
        urlChars = ( alnum | "_" | "~" | "." | "-" ) $store;

        main := ( urlChars | percentEncode | spaceChar )* %{ return true, decoded };

        write init;
        write exec;
	}%%

	return false, nil
}
