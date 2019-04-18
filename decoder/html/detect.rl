package html

func (d *decoder) Detect(data []byte) bool {

	%% machine detect_html;
	%% write data;

	cs, p, pe, eof := 0, 0, len(data), len(data)
	_ = eof

	%%{
        entityName = ( alpha+ );
        entityHex = ( "#x" digit+ );
        entityNumber = ( "#" digit+ );
        main := any* "&" ( entityName | entityHex | entityNumber ) ";" @{ return true };

        write init;
        write exec;
	}%%

	return false
}
