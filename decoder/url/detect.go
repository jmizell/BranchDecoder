
//line ./decoder/url/detect.rl:1
package url

func (d *decoder) Detect(data []byte) bool {

	
//line ./decoder/url/detect.rl:6
	
//line ./decoder/url/detect.go:11
var _detect_url_eof_actions []byte = []byte{
	0, 0, 0, 1, 2, 
}

const detect_url_start int = 3
const detect_url_first_final int = 3
const detect_url_error int = 0

const detect_url_en_main int = 3


//line ./decoder/url/detect.rl:7

	cs, p, pe, eof := 0, 0, len(data), len(data)
    detectedEscapes := false
    match := false
	_ = eof

	
//line ./decoder/url/detect.go:31
	{
	cs = detect_url_start
	}

//line ./decoder/url/detect.go:36
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
			goto tr3;
		case 43:
			goto tr2;
		case 95:
			goto tr4;
		case 126:
			goto tr4;
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr4;
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr4;
				}
			case data[p] >= 65:
				goto tr4;
			}
		default:
			goto tr4;
		}
		goto tr1;
	case 0:
		goto _out
	case 1:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr0;
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr0;
			}
		default:
			goto tr0;
		}
		goto tr1;
	case 2:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr2;
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr2;
			}
		default:
			goto tr2;
		}
		goto tr1;
	case 4:
		switch data[p] {
		case 37:
			goto tr5;
		case 43:
			goto tr6;
		case 95:
			goto tr7;
		case 126:
			goto tr7;
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr7;
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr7;
				}
			case data[p] >= 65:
				goto tr7;
			}
		default:
			goto tr7;
		}
		goto tr1;
	}

	tr1: cs = 0; goto _again
	tr3: cs = 1; goto _again
	tr5: cs = 1; goto f2
	tr0: cs = 2; goto _again
	tr4: cs = 3; goto _again
	tr7: cs = 3; goto f2
	tr2: cs = 4; goto _again
	tr6: cs = 4; goto f2

f2:
//line ./decoder/url/detect.rl:14
 detectedEscapes = true 
	goto _again

_again:
	if cs == 0 {
		goto _out
	}
	if p++; p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		switch _detect_url_eof_actions[cs] {
		case 1:
//line ./decoder/url/detect.rl:16
 match = true 
		case 2:
//line ./decoder/url/detect.rl:14
 detectedEscapes = true 
//line ./decoder/url/detect.rl:16
 match = true 
//line ./decoder/url/detect.go:168
		}
	}

	_out: {}
	}

//line ./decoder/url/detect.rl:20


	return match && detectedEscapes
}
