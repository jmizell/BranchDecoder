package url

func (d *decoder) Detect(data []byte) bool {

	%% machine detect_url;
	%% write data;

	cs, p, pe, eof := 0, 0, len(data), len(data)
    detectedEscapes := false
    match := false
	_ = eof

	%%{
        escapeChars = ( "+" | "%" xdigit xdigit ) %{ detectedEscapes = true };
        urlChars = ( alnum | "_" | "~" | "." | "-" | escapeChars );
        main := urlChars* %{ match = true };

        write init;
        write exec;
	}%%

	return match && detectedEscapes
}
