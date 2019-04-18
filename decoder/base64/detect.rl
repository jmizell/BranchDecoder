package base64

func (d *decoder) Detect(data []byte) bool {

	%% machine detect_base64;
	%% write data;

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	%%{
        main := ( alnum | "=" | "+" | "/" )+ %{ return true } | ( alnum | "=" | "-" | "_" )+  %{ return true };

        write init;
        write exec;
	}%%

	return false
}
