//line ./decoder/url/decode.rl:1
package url

func hexToByte(char byte) byte {

	if 'a' <= char && char <= 'f' {
		return char - 'a' + 10
	}
	if 'A' <= char && char <= 'F' {
		return char - 'A' + 10
	}
	if '0' <= char && char <= '9' {
		return char - '0'
	}

	return 0
}

func (d *decoder) Decode(data []byte) (match bool, decoded []byte) {

//line ./decoder/url/decode.rl:15

//line ./decoder/url/decode.go:20
	var _decode_url_eof_actions []byte = []byte{
		0, 0, 0, 2, 5,
	}

	const decode_url_start int = 3
	const decode_url_first_final int = 3
	const decode_url_error int = 0

	const decode_url_en_main int = 3

//line ./decoder/url/decode.rl:16

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

			for _, d := range []int{1, 2} {
				if !('0' <= data[i+d] && data[i+d] <= '9' ||
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

//line ./decoder/url/decode.go:76
	{
		cs = decode_url_start
	}

//line ./decoder/url/decode.go:81
	{
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		switch cs {
		case 3:
			switch data[p] {
			case 37:
				goto tr3
			case 43:
				goto tr4
			case 95:
				goto tr5
			case 126:
				goto tr5
			}
			switch {
			case data[p] < 48:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr5
				}
			case data[p] > 57:
				switch {
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr5
					}
				case data[p] >= 65:
					goto tr5
				}
			default:
				goto tr5
			}
			goto tr1
		case 0:
			goto _out
		case 1:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr0
				}
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto tr0
				}
			default:
				goto tr0
			}
			goto tr1
		case 2:
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr2
				}
			case data[p] > 70:
				if 97 <= data[p] && data[p] <= 102 {
					goto tr2
				}
			default:
				goto tr2
			}
			goto tr1
		case 4:
			switch data[p] {
			case 37:
				goto tr6
			case 43:
				goto tr7
			case 95:
				goto tr8
			case 126:
				goto tr8
			}
			switch {
			case data[p] < 48:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr8
				}
			case data[p] > 57:
				switch {
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr8
					}
				case data[p] >= 65:
					goto tr8
				}
			default:
				goto tr8
			}
			goto tr1
		}

	tr1:
		cs = 0
		goto _again
	tr3:
		cs = 1
		goto _again
	tr6:
		cs = 1
		goto f5
	tr0:
		cs = 2
		goto f0
	tr4:
		cs = 3
		goto f2
	tr5:
		cs = 3
		goto f3
	tr7:
		cs = 3
		goto f6
	tr8:
		cs = 3
		goto f7
	tr2:
		cs = 4
		goto _again

	f0:
//line ./decoder/url/decode.rl:59
		mark = p
		goto _again
	f5:
//line ./decoder/url/decode.rl:60
		decoded[j] = hexToByte(data[mark:p][0])<<4 | hexToByte(data[mark:p][1])
		j++
		goto _again
	f2:
//line ./decoder/url/decode.rl:61
		decoded[j] = ' '
		j++
		goto _again
	f3:
//line ./decoder/url/decode.rl:62
		decoded[j] = data[p]
		j++
		goto _again
	f6:
//line ./decoder/url/decode.rl:60
		decoded[j] = hexToByte(data[mark:p][0])<<4 | hexToByte(data[mark:p][1])
		j++
//line ./decoder/url/decode.rl:61
		decoded[j] = ' '
		j++
		goto _again
	f7:
//line ./decoder/url/decode.rl:60
		decoded[j] = hexToByte(data[mark:p][0])<<4 | hexToByte(data[mark:p][1])
		j++
//line ./decoder/url/decode.rl:62
		decoded[j] = data[p]
		j++
		goto _again

	_again:
		if cs == 0 {
			goto _out
		}
		if p++; p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
		if p == eof {
			switch _decode_url_eof_actions[cs] {
			case 2:
//line ./decoder/url/decode.rl:68
				return true, decoded
			case 5:
//line ./decoder/url/decode.rl:60
				decoded[j] = hexToByte(data[mark:p][0])<<4 | hexToByte(data[mark:p][1])
				j++
//line ./decoder/url/decode.rl:68
				return true, decoded
//line ./decoder/url/decode.go:238
			}
		}

	_out:
		{
		}
	}

//line ./decoder/url/decode.rl:72

	return false, nil
}
