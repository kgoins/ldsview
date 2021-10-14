package searcher

import "github.com/kgoins/ldifparser"

const LDIFSEARCHER_DEFAULT_BUFF_SIZE = 128000

type LdifSearcherConf struct {
	BufferSize int
}

func NewLdifSearcherConf() LdifSearcherConf {
	return LdifSearcherConf{
		BufferSize: LDIFSEARCHER_DEFAULT_BUFF_SIZE,
	}
}

func buildLdifReaderConf(conf LdifSearcherConf) ldifparser.ReaderConf {
	readerConf := ldifparser.NewReaderConf()

	readerConf.ScannerBufferSize = conf.BufferSize

	return readerConf
}
