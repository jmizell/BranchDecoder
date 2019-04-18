//line ./decoder/html/detect.rl:1
package html

func (d *decoder) Detect(data []byte) bool {

//line ./decoder/html/detect.rl:6

//line ./decoder/html/detect.go:11
	const detect_html_start int = 0
	const detect_html_first_final int = 6
	const detect_html_error int = -1

	const detect_html_en_main int = 0

//line ./decoder/html/detect.rl:7

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

//line ./decoder/html/detect.go:25
	{
		cs = detect_html_start
	}

//line ./decoder/html/detect.go:30
	{
		if p == pe {
			goto _test_eof
		}
	_resume:
		switch cs {
		case 0:
			if data[p] == 38 {
				goto tr1
			}
			goto tr0
		case 1:
			switch data[p] {
			case 35:
				goto tr2
			case 38:
				goto tr1
			}
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			case data[p] >= 65:
				goto tr3
			}
			goto tr0
		case 2:
			switch data[p] {
			case 38:
				goto tr1
			case 120:
				goto tr5
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
			goto tr0
		case 3:
			switch data[p] {
			case 38:
				goto tr1
			case 59:
				goto tr6
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
			goto tr0
		case 6:
			if data[p] == 38 {
				goto tr1
			}
			goto tr0
		case 4:
			if data[p] == 38 {
				goto tr1
			}
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
			goto tr0
		case 5:
			switch data[p] {
			case 38:
				goto tr1
			case 59:
				goto tr6
			}
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			case data[p] >= 65:
				goto tr3
			}
			goto tr0
		}

	tr0:
		cs = 0
		goto _again
	tr1:
		cs = 1
		goto _again
	tr2:
		cs = 2
		goto _again
	tr4:
		cs = 3
		goto _again
	tr5:
		cs = 4
		goto _again
	tr3:
		cs = 5
		goto _again
	tr6:
		cs = 6
		goto f0

	f0:
//line ./decoder/html/detect.rl:15
		return true
		goto _again

	_again:
		if p++; p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
	}

//line ./decoder/html/detect.rl:19

	return false
}
