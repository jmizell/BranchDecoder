//line ./decoder/base64/detect.rl:1
package base64

func (d *decoder) Detect(data []byte) bool {

//line ./decoder/base64/detect.rl:6

//line ./decoder/base64/detect.go:11
	var _detect_base64_eof_actions []byte = []byte{
		0, 0, 1, 2, 3,
	}

	const detect_base64_start int = 1
	const detect_base64_first_final int = 2
	const detect_base64_error int = 0

	const detect_base64_en_main int = 1

//line ./decoder/base64/detect.rl:7

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

//line ./decoder/base64/detect.go:29
	{
		cs = detect_base64_start
	}

//line ./decoder/base64/detect.go:34
	{
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		switch cs {
		case 1:
			switch data[p] {
			case 43:
				goto tr0
			case 45:
				goto tr2
			case 47:
				goto tr0
			case 61:
				goto tr3
			case 95:
				goto tr2
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr3
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			default:
				goto tr3
			}
			goto tr1
		case 0:
			goto _out
		case 2:
			switch data[p] {
			case 43:
				goto tr0
			case 61:
				goto tr0
			}
			switch {
			case data[p] < 65:
				if 47 <= data[p] && data[p] <= 57 {
					goto tr0
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr0
				}
			default:
				goto tr0
			}
			goto tr1
		case 3:
			switch data[p] {
			case 45:
				goto tr2
			case 61:
				goto tr2
			case 95:
				goto tr2
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr2
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr2
				}
			default:
				goto tr2
			}
			goto tr1
		case 4:
			switch data[p] {
			case 43:
				goto tr0
			case 45:
				goto tr2
			case 47:
				goto tr0
			case 61:
				goto tr3
			case 95:
				goto tr2
			}
			switch {
			case data[p] < 65:
				if 48 <= data[p] && data[p] <= 57 {
					goto tr3
				}
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			default:
				goto tr3
			}
			goto tr1
		}

	tr1:
		cs = 0
		goto _again
	tr0:
		cs = 2
		goto _again
	tr2:
		cs = 3
		goto _again
	tr3:
		cs = 4
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
			switch _detect_base64_eof_actions[cs] {
			case 1:
//line ./decoder/base64/detect.rl:12
				return true
			case 2:
//line ./decoder/base64/detect.rl:12
				return true
			case 3:
//line ./decoder/base64/detect.rl:12
				return true
//line ./decoder/base64/detect.rl:12
				return true
//line ./decoder/base64/detect.go:168
			}
		}

	_out:
		{
		}
	}

//line ./decoder/base64/detect.rl:16

	return false
}
