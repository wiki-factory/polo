package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func templates_archives_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4c, 0x8f,
		0x41, 0x4e, 0xc4, 0x30, 0x0c, 0x45, 0xf7, 0x3d, 0x85, 0x15, 0xcd, 0x12,
		0x52, 0xb1, 0x45, 0x69, 0xa4, 0x11, 0xdc, 0x00, 0xf6, 0x28, 0x6a, 0xdc,
		0x26, 0x92, 0x49, 0xa5, 0xc4, 0xb0, 0xc0, 0xe4, 0xee, 0x78, 0x28, 0xc3,
		0x74, 0x95, 0xd8, 0xcf, 0xff, 0xdb, 0x5f, 0x24, 0xe2, 0x92, 0x0b, 0x82,
		0xe1, 0xcc, 0x84, 0xa6, 0xf7, 0xe1, 0x5c, 0xe7, 0x94, 0x3f, 0xb1, 0xc1,
		0x37, 0x88, 0xd8, 0xa7, 0xad, 0x2c, 0x79, 0xb5, 0xaf, 0x17, 0xaa, 0x50,
		0x04, 0x4b, 0xd4, 0x77, 0xb8, 0x09, 0xe7, 0xad, 0x30, 0x16, 0xbe, 0x48,
		0x5d, 0x7a, 0xf0, 0x57, 0xb9, 0x1b, 0xb5, 0x18, 0x5c, 0x24, 0x98, 0x29,
		0xb4, 0x36, 0x99, 0x48, 0xf7, 0x69, 0xab, 0xf9, 0x4b, 0xc7, 0x03, 0x19,
		0x3f, 0x80, 0xda, 0xd7, 0x50, 0x56, 0x84, 0xd3, 0xdb, 0x1d, 0x9c, 0x42,
		0xe5, 0x3c, 0x13, 0xc2, 0xe3, 0x04, 0xf6, 0xbc, 0xff, 0x9b, 0x5a, 0x02,
		0xb8, 0xc8, 0x5e, 0xe4, 0xca, 0xed, 0x73, 0x60, 0x3d, 0xc4, 0x8d, 0xda,
		0xfd, 0x85, 0xd1, 0xbb, 0x00, 0xa9, 0xe2, 0x32, 0x99, 0xf1, 0x30, 0xf6,
		0x42, 0x1f, 0x6b, 0xef, 0x36, 0xf1, 0xbb, 0xee, 0x3a, 0xf4, 0xff, 0x82,
		0xb8, 0x31, 0x78, 0xf5, 0x88, 0xfb, 0x19, 0x7b, 0x26, 0xad, 0xc9, 0xff,
		0x27, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x61, 0x4f, 0x8b, 0x01, 0x1a,
		0x01, 0x00, 0x00,
	},
		"templates/archives.html",
	)
}

func templates_article_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x54,
		0x51, 0x4f, 0xe3, 0x38, 0x10, 0x7e, 0xef, 0xaf, 0x18, 0x45, 0x48, 0x49,
		0xc5, 0x11, 0xdf, 0xf1, 0x08, 0x69, 0x4e, 0x5c, 0x39, 0xe9, 0x6e, 0x1f,
		0x96, 0x95, 0xe0, 0x6d, 0xb5, 0x42, 0x4e, 0xe2, 0x24, 0x06, 0xd7, 0x0e,
		0xb6, 0xd3, 0x6e, 0x28, 0xfd, 0xef, 0x3b, 0x8e, 0x13, 0x9a, 0x42, 0x11,
		0xfb, 0xd2, 0x34, 0x33, 0xf3, 0x7d, 0xf3, 0x79, 0xe6, 0x73, 0xb6, 0xdb,
		0x82, 0x95, 0x5c, 0x32, 0x08, 0x2c, 0xb7, 0x82, 0x05, 0xbb, 0xdd, 0x6c,
		0xbb, 0x8d, 0xaf, 0xb4, 0xe5, 0xb9, 0x60, 0xf1, 0x9d, 0x8b, 0xed, 0x76,
		0xf0, 0x02, 0x18, 0x5c, 0x2a, 0x59, 0xf2, 0x6a, 0x8c, 0x61, 0x19, 0x93,
		0x05, 0x3e, 0x67, 0x7b, 0x8a, 0x5c, 0x49, 0xcb, 0xa4, 0x75, 0x24, 0x49,
		0xfd, 0x57, 0xfa, 0x9e, 0x28, 0x21, 0x18, 0x9e, 0xcd, 0x92, 0x26, 0x45,
		0x14, 0x2f, 0x61, 0xe4, 0xbc, 0x6a, 0x6d, 0xad, 0x34, 0x38, 0x9c, 0x69,
		0xa8, 0x84, 0x5c, 0x50, 0x63, 0x16, 0x81, 0xa0, 0x19, 0x13, 0xd0, 0xff,
		0x9e, 0x61, 0x0f, 0xda, 0x0a, 0x1b, 0xa4, 0x59, 0x37, 0x11, 0xe3, 0x81,
		0x8e, 0xd8, 0xe1, 0xd2, 0xa9, 0x28, 0x47, 0x3f, 0xb6, 0xbf, 0xa6, 0x96,
		0xfd, 0x16, 0xf9, 0x44, 0xb2, 0xc7, 0x7c, 0x42, 0xbc, 0xc4, 0xa2, 0x4a,
		0xe9, 0xce, 0x91, 0xd3, 0x63, 0xcc, 0xa6, 0xcd, 0x73, 0x66, 0x4c, 0x00,
		0xb5, 0x66, 0xe5, 0x22, 0x20, 0xf9, 0x00, 0x20, 0x93, 0x4e, 0x7b, 0x92,
		0xb8, 0xb6, 0x2b, 0x71, 0xa0, 0x62, 0x9f, 0x4b, 0x08, 0xfd, 0x50, 0xc6,
		0x1d, 0xad, 0x4c, 0xbf, 0x12, 0x4d, 0x65, 0xc5, 0xe0, 0xe4, 0xfe, 0x0f,
		0x38, 0xb1, 0xb4, 0x82, 0x8b, 0xc5, 0xbb, 0x9a, 0xe3, 0x32, 0xb9, 0x2c,
		0xd5, 0xab, 0x46, 0x44, 0xa2, 0x3c, 0x47, 0x30, 0x51, 0xe4, 0x5f, 0x0f,
		0x44, 0x8c, 0xcf, 0x84, 0xe0, 0x3e, 0xa7, 0xb6, 0xf9, 0x0f, 0x31, 0x93,
		0x19, 0x79, 0x57, 0xf4, 0x80, 0xc9, 0xce, 0xaf, 0xb9, 0x79, 0x6a, 0xcd,
		0x2d, 0xc7, 0x1c, 0x5d, 0xf5, 0xdb, 0xa9, 0xcf, 0xd3, 0xa5, 0x5a, 0xad,
		0xb0, 0xd6, 0xa0, 0x51, 0xce, 0xd3, 0x59, 0x52, 0xf0, 0x35, 0xf0, 0x62,
		0x11, 0x14, 0x7d, 0xed, 0xbd, 0x45, 0x81, 0xb4, 0x08, 0xd2, 0x84, 0x60,
		0x02, 0xd3, 0x26, 0xd7, 0xbc, 0xb1, 0x60, 0xbb, 0x86, 0x2d, 0x02, 0xcb,
		0x7e, 0x5a, 0xf2, 0x40, 0xd7, 0xd4, 0x47, 0x83, 0x74, 0x06, 0xb0, 0xa6,
		0x1a, 0x06, 0xac, 0x41, 0x9f, 0x58, 0xd7, 0x09, 0x16, 0x10, 0xee, 0x0d,
		0xf4, 0x56, 0x45, 0x78, 0x89, 0x30, 0x42, 0xe0, 0xee, 0xe6, 0xfa, 0x06,
		0x22, 0x5a, 0x29, 0xf9, 0x4c, 0x05, 0x7b, 0xd6, 0x6a, 0x7e, 0x01, 0x28,
		0xae, 0xa1, 0x96, 0x67, 0x5c, 0x70, 0xdb, 0xc1, 0x86, 0xdb, 0x1a, 0x56,
		0x1d, 0x28, 0x51, 0x40, 0x26, 0x54, 0x05, 0xd4, 0x9f, 0xd7, 0x78, 0x82,
		0xff, 0xa1, 0x50, 0x32, 0xb4, 0xf0, 0x28, 0xd5, 0x06, 0xf0, 0xd8, 0xdc,
		0x22, 0x42, 0x08, 0xd8, 0x28, 0xfd, 0x08, 0x25, 0x9a, 0x5d, 0xd9, 0x9a,
		0x69, 0x68, 0x0d, 0xd3, 0xe6, 0x50, 0x29, 0x2f, 0x70, 0x02, 0xbc, 0xe4,
		0x98, 0xf5, 0x52, 0xc7, 0x41, 0xde, 0x8a, 0x76, 0xdc, 0x08, 0xca, 0x44,
		0x50, 0x54, 0xb6, 0x32, 0xb7, 0x5c, 0xc9, 0x68, 0x0e, 0x5b, 0x7c, 0x1f,
		0x68, 0xcc, 0x13, 0x02, 0x0b, 0x95, 0xb7, 0x6e, 0x94, 0x71, 0x8e, 0x33,
		0xb3, 0xec, 0x5f, 0xc1, 0xdc, 0x5b, 0x14, 0xfa, 0xe9, 0x84, 0xf3, 0x4b,
		0x57, 0x17, 0xbb, 0xd1, 0xb9, 0x2e, 0x6f, 0x86, 0x17, 0xfa, 0x2c, 0x35,
		0x9d, 0xcc, 0x31, 0x6d, 0x75, 0xcb, 0x2e, 0x7b, 0x7e, 0x17, 0x35, 0xda,
		0xc5, 0x42, 0x42, 0x42, 0x38, 0x7d, 0x3f, 0xdd, 0x53, 0x08, 0x63, 0x1f,
		0x8c, 0x73, 0xb5, 0x22, 0x6c, 0x95, 0xb1, 0x22, 0x7e, 0x30, 0xa1, 0xc7,
		0x47, 0xaf, 0xb2, 0x2a, 0x66, 0x07, 0x4d, 0xe6, 0x9f, 0x0e, 0xbd, 0xf9,
		0x15, 0xc1, 0x51, 0x58, 0xe3, 0x7e, 0xc3, 0xf9, 0xf7, 0x3f, 0x7f, 0xc0,
		0xcb, 0x0b, 0x7c, 0x52, 0x9b, 0xa9, 0xa2, 0xeb, 0x6b, 0xe7, 0x31, 0x6d,
		0x1a, 0xb4, 0xe2, 0xb2, 0xe6, 0xa2, 0x88, 0x50, 0xe2, 0xdc, 0x35, 0xdb,
		0xcd, 0x23, 0x7c, 0xe2, 0xfd, 0xed, 0x4f, 0x84, 0x6e, 0x91, 0x6a, 0xf8,
		0xfb, 0x4d, 0x30, 0x6a, 0x18, 0xe0, 0xc6, 0x33, 0xc1, 0xe0, 0x0b, 0x9e,
		0xfa, 0x76, 0x30, 0x92, 0x82, 0x35, 0x67, 0x1b, 0xc0, 0xcd, 0x00, 0xde,
		0x13, 0x7f, 0x1d, 0x6a, 0x6b, 0x9b, 0x0b, 0x42, 0x26, 0x67, 0xfa, 0x1b,
		0xe3, 0xf7, 0x23, 0x5b, 0x90, 0xe6, 0x83, 0x67, 0xa1, 0x51, 0x1b, 0xa6,
		0x19, 0xba, 0xa1, 0x03, 0x6f, 0xab, 0xd8, 0x5d, 0x97, 0x84, 0xbc, 0x36,
		0x9e, 0x7d, 0x4c, 0x1a, 0x8c, 0xb7, 0x12, 0xe5, 0x9f, 0x65, 0x5a, 0x70,
		0xf9, 0x78, 0x9c, 0xf9, 0xf0, 0x1b, 0xa6, 0x2a, 0x75, 0xe6, 0x49, 0x82,
		0xd4, 0xf7, 0x1c, 0x3e, 0x58, 0x47, 0x2f, 0xea, 0xaf, 0x00, 0x00, 0x00,
		0xff, 0xff, 0x81, 0x25, 0xb5, 0x38, 0xed, 0x05, 0x00, 0x00,
	},
		"templates/article.html",
	)
}

func templates_atom_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x6c, 0x92,
		0x5d, 0x4b, 0xc3, 0x30, 0x18, 0x85, 0xef, 0xfb, 0x2b, 0x42, 0xd8, 0xa5,
		0x26, 0x53, 0x11, 0x64, 0xa4, 0x1d, 0x43, 0x11, 0x2f, 0xbc, 0x72, 0xee,
		0x5a, 0xc2, 0x9a, 0xb6, 0xc1, 0x7c, 0x8c, 0x34, 0xb5, 0x4a, 0xc9, 0x7f,
		0xb7, 0x49, 0x97, 0x7e, 0x8c, 0xdd, 0xa5, 0x39, 0xe7, 0xbc, 0xef, 0x73,
		0x42, 0xc9, 0xf6, 0x57, 0x0a, 0xf0, 0xc3, 0x4c, 0xcd, 0xb5, 0x4a, 0xe1,
		0x1d, 0x5a, 0x43, 0xc0, 0xd4, 0x51, 0xe7, 0x5c, 0x95, 0x29, 0x6c, 0x6c,
		0x71, 0xfb, 0x04, 0xb7, 0x59, 0x92, 0x90, 0x82, 0xb1, 0x1c, 0xf4, 0x5e,
		0x55, 0xa7, 0xb0, 0xb2, 0xf6, 0xb4, 0xc1, 0xb8, 0x6d, 0x5b, 0xd4, 0x3e,
		0x20, 0x6d, 0x4a, 0x7c, 0xbf, 0x5e, 0x3f, 0xe2, 0x9d, 0xd5, 0x12, 0x66,
		0x09, 0x00, 0xc4, 0x72, 0x2b, 0x58, 0xd6, 0x75, 0xe8, 0x59, 0xab, 0x82,
		0x97, 0xe8, 0xd3, 0x7f, 0x3b, 0x47, 0xf0, 0x20, 0x78, 0x8b, 0xe0, 0xea,
		0x1b, 0x54, 0x86, 0x15, 0x29, 0x9c, 0x7c, 0x87, 0x8f, 0x77, 0xe7, 0x20,
		0xc0, 0xc1, 0xd1, 0x9c, 0x72, 0x6a, 0x59, 0xee, 0xc7, 0x1c, 0x86, 0xa3,
		0x9f, 0x10, 0x6f, 0x93, 0xde, 0xd2, 0x75, 0xbc, 0x00, 0xe8, 0xb5, 0x27,
		0xdb, 0x19, 0xcb, 0x8f, 0x82, 0xd5, 0xce, 0x85, 0x6b, 0x43, 0x55, 0xc9,
		0xc0, 0xea, 0xeb, 0x06, 0xac, 0xe8, 0xa0, 0x80, 0x4d, 0x7a, 0xc5, 0x49,
		0x98, 0xb2, 0xe6, 0xcf, 0x6f, 0x9b, 0x41, 0xc7, 0xc8, 0x15, 0xea, 0x4b,
		0xee, 0xd5, 0x02, 0x1c, 0xcf, 0xb2, 0x7b, 0xd1, 0x94, 0xce, 0xa1, 0xca,
		0x4a, 0x71, 0xee, 0x73, 0xc6, 0x1d, 0x1d, 0x2f, 0x7d, 0x8d, 0x00, 0xb1,
		0xa8, 0x7a, 0x21, 0xcf, 0xfa, 0x0e, 0x13, 0x98, 0xca, 0x63, 0xa8, 0x6e,
		0xa4, 0xa4, 0x91, 0xde, 0x6b, 0x63, 0xf6, 0xad, 0xdf, 0x3a, 0x2d, 0xda,
		0x0f, 0xbe, 0x18, 0xc3, 0x8b, 0x1c, 0xa1, 0x8d, 0xad, 0xb4, 0x89, 0x43,
		0x88, 0xa2, 0x32, 0xbc, 0x41, 0x2c, 0xb6, 0x0b, 0xb2, 0x07, 0x09, 0xca,
		0x79, 0xc4, 0x14, 0x22, 0x78, 0x7c, 0xc2, 0x09, 0x2e, 0x9e, 0x08, 0xf6,
		0x7f, 0x4d, 0x96, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x01, 0xd4,
		0xd7, 0x65, 0x02, 0x00, 0x00,
	},
		"templates/atom.xml",
	)
}

func templates_category_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x54, 0x8e,
		0x41, 0x8b, 0xc3, 0x20, 0x14, 0x84, 0xef, 0xfe, 0x8a, 0x87, 0xe4, 0xb8,
		0x28, 0x7b, 0x5d, 0x5e, 0x84, 0x6d, 0x7f, 0x42, 0x7b, 0x2f, 0x92, 0xbc,
		0x18, 0xc1, 0x1a, 0xb0, 0xe6, 0x10, 0xac, 0xff, 0xbd, 0x9a, 0x34, 0x34,
		0x3d, 0xe9, 0x0c, 0xf3, 0xcd, 0x9b, 0x94, 0x7a, 0x1a, 0xac, 0x27, 0xe0,
		0xd1, 0x46, 0x47, 0x3c, 0x67, 0x96, 0x92, 0x38, 0xeb, 0x48, 0x66, 0x0a,
		0x4b, 0xce, 0xf0, 0x84, 0xaa, 0x27, 0x3f, 0x58, 0x23, 0xae, 0x35, 0xb2,
		0x26, 0xc8, 0xf7, 0xe5, 0x65, 0x1f, 0xba, 0x9b, 0x7c, 0x24, 0x1f, 0x2b,
		0x8f, 0xe3, 0xaf, 0xfa, 0xea, 0x40, 0x59, 0x1c, 0x86, 0xb3, 0x53, 0x0c,
		0x4a, 0x5b, 0xd0, 0xde, 0x10, 0x34, 0xb7, 0x1f, 0x68, 0x74, 0x88, 0xb6,
		0x73, 0x04, 0x7f, 0x2d, 0x88, 0xff, 0xed, 0xff, 0x38, 0x2d, 0x3b, 0x09,
		0x87, 0x8e, 0x42, 0xa2, 0xb3, 0x0a, 0x35, 0x8c, 0x81, 0x86, 0x96, 0xcb,
		0x94, 0x76, 0x5a, 0x5c, 0xdc, 0x6c, 0x72, 0x16, 0x63, 0xbc, 0x3b, 0xae,
		0x0e, 0xfe, 0x7b, 0x2e, 0x4a, 0xad, 0x50, 0x16, 0x78, 0xbd, 0xbe, 0x2d,
		0x47, 0x59, 0xd7, 0xec, 0xea, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xbc,
		0x07, 0xa3, 0x05, 0x01, 0x00, 0x00,
	},
		"templates/category.html",
	)
}

func templates_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x55,
		0xdf, 0x6f, 0x9b, 0x30, 0x10, 0x7e, 0xe7, 0xaf, 0x70, 0x59, 0xb4, 0x87,
		0x6a, 0x05, 0xe5, 0x75, 0x23, 0x48, 0x5d, 0x27, 0x6d, 0x0f, 0x53, 0x56,
		0x69, 0x7d, 0xaf, 0x1c, 0x38, 0xc0, 0x92, 0x63, 0x12, 0xdb, 0x44, 0x8d,
		0x10, 0xff, 0xfb, 0xce, 0x3f, 0x08, 0x26, 0x49, 0xd7, 0xaa, 0x95, 0x8a,
		0xb1, 0x7d, 0xf7, 0xdd, 0x77, 0x77, 0xdf, 0x91, 0xbe, 0x2f, 0xa1, 0x62,
		0x02, 0x48, 0xac, 0x99, 0xe6, 0x10, 0x0f, 0x43, 0xd4, 0xf7, 0xc9, 0x43,
		0x2b, 0x2a, 0x56, 0x27, 0x4f, 0xe6, 0xc8, 0x9e, 0x80, 0x28, 0x71, 0x8d,
		0x26, 0xeb, 0xa2, 0x15, 0x1a, 0x84, 0x76, 0xf6, 0xac, 0x22, 0xc9, 0xbd,
		0xd4, 0xac, 0xe0, 0xa0, 0xbe, 0x1f, 0x1f, 0x69, 0x0d, 0x24, 0x31, 0xcf,
		0x75, 0xb7, 0xdd, 0x80, 0x74, 0x8e, 0x92, 0x0a, 0x3c, 0x5e, 0x3c, 0x7f,
		0x21, 0x0b, 0xea, 0x4c, 0xc9, 0xd7, 0xd5, 0x1b, 0x6e, 0x99, 0xb7, 0xcc,
		0xa3, 0xac, 0x59, 0xe6, 0x19, 0x25, 0x8d, 0x84, 0x6a, 0x15, 0xf7, 0xfd,
		0x08, 0x91, 0xfc, 0xe5, 0x5d, 0x3d, 0x0c, 0x49, 0xa3, 0xb7, 0x3c, 0xce,
		0x83, 0x73, 0xcf, 0x3c, 0x4b, 0x69, 0x9e, 0xa5, 0xe8, 0x1b, 0x45, 0xd9,
		0x2e, 0x77, 0x4c, 0x5b, 0x49, 0x16, 0x63, 0x82, 0xf7, 0x9d, 0x6e, 0xcc,
		0x7e, 0x74, 0x73, 0x7b, 0x13, 0x59, 0xed, 0xa8, 0x20, 0x05, 0xa7, 0x4a,
		0xad, 0x62, 0x4e, 0x37, 0xc0, 0x89, 0x7d, 0xde, 0x61, 0xfe, 0xb4, 0xe3,
		0x3a, 0xce, 0x37, 0x47, 0x62, 0xe1, 0x2e, 0x7c, 0x03, 0x16, 0xd3, 0x11,
		0x70, 0x05, 0xf6, 0x6a, 0x1e, 0xd9, 0x5e, 0x99, 0xd2, 0x66, 0xa9, 0x09,
		0x98, 0x87, 0x95, 0x0e, 0xb1, 0x7f, 0x50, 0x0d, 0xef, 0x62, 0x15, 0x04,
		0x77, 0x3e, 0x6f, 0x00, 0x3f, 0xa0, 0x51, 0xdd, 0xca, 0xa3, 0x2d, 0xf6,
		0x35, 0x64, 0xd5, 0x15, 0x05, 0x28, 0x15, 0xfb, 0xd2, 0xa7, 0x85, 0x77,
		0x48, 0x83, 0x48, 0x13, 0xc8, 0x65, 0x23, 0xa6, 0x3b, 0xd3, 0x8b, 0xd7,
		0x68, 0x3c, 0xd1, 0x5a, 0x59, 0x25, 0x05, 0x2a, 0xd1, 0xb4, 0x36, 0x0a,
		0x39, 0xb7, 0xb9, 0x4e, 0x93, 0x89, 0xaa, 0x3d, 0x71, 0x44, 0x4f, 0x43,
		0x0f, 0x97, 0x90, 0x91, 0xdd, 0xce, 0x48, 0x8c, 0x6b, 0x96, 0xa2, 0x36,
		0xa2, 0x80, 0xf4, 0x2f, 0xf4, 0x09, 0x6a, 0xe4, 0xa4, 0x6e, 0x0c, 0x77,
		0x63, 0x6c, 0x0d, 0x2f, 0xfa, 0x4e, 0xb2, 0xba, 0xc1, 0x9a, 0x4f, 0x8c,
		0x36, 0x5a, 0x10, 0xfc, 0x3f, 0xb5, 0xe3, 0xbf, 0x7a, 0xfd, 0x54, 0x32,
		0xb5, 0xef, 0xd4, 0xb3, 0x46, 0x23, 0x5a, 0xc6, 0xf9, 0x6f, 0xa0, 0x07,
		0x20, 0x88, 0xd5, 0x6e, 0xb7, 0x18, 0xee, 0xc6, 0x69, 0x17, 0x99, 0xe1,
		0xcb, 0x34, 0x06, 0x72, 0x56, 0x43, 0x27, 0x2b, 0x3b, 0x1d, 0xeb, 0x56,
		0x37, 0x4c, 0xd4, 0x44, 0xb7, 0x44, 0x01, 0x90, 0x06, 0x24, 0xdc, 0x38,
		0xe9, 0x9f, 0x3b, 0x99, 0x69, 0x75, 0x23, 0xf6, 0xa7, 0x32, 0xe3, 0x66,
		0xab, 0x5a, 0xb2, 0xc3, 0x2c, 0xb7, 0x02, 0x29, 0x80, 0xc4, 0xe4, 0x3a,
		0x3e, 0x9e, 0xef, 0x68, 0xcd, 0x04, 0xd5, 0xac, 0x15, 0x31, 0x96, 0x8b,
		0xe0, 0x5f, 0xc6, 0xd9, 0x78, 0x69, 0xf7, 0x16, 0x1b, 0xf6, 0x38, 0x5f,
		0xd3, 0x18, 0x93, 0x25, 0xc2, 0x9b, 0x4b, 0x4c, 0x97, 0x6e, 0x38, 0x94,
		0xde, 0xd2, 0xd1, 0x31, 0xef, 0x88, 0x66, 0xc1, 0x82, 0xf1, 0x4e, 0x7e,
		0x82, 0x7e, 0x94, 0x70, 0x60, 0x6d, 0xa7, 0x0c, 0x94, 0x29, 0xdb, 0x0c,
		0x75, 0x18, 0xe2, 0xfc, 0x33, 0xa7, 0xfb, 0xae, 0xfd, 0x66, 0x5b, 0x6a,
		0x01, 0x52, 0xce, 0x3c, 0xb1, 0x99, 0x90, 0x76, 0xe6, 0xc3, 0xe2, 0xbe,
		0x35, 0x92, 0x1e, 0xa7, 0x9c, 0xc7, 0x0c, 0x4e, 0xac, 0xad, 0xe1, 0x3c,
		0x8a, 0xcf, 0x8e, 0x16, 0x9a, 0x1d, 0x20, 0xf6, 0xac, 0xcf, 0xf8, 0xa6,
		0x4c, 0x94, 0xf0, 0x62, 0x51, 0xf0, 0xcb, 0xe8, 0x50, 0x96, 0x76, 0xd8,
		0xcd, 0xeb, 0x69, 0xc4, 0xbd, 0x12, 0xad, 0xaf, 0x61, 0xe8, 0x6f, 0x3d,
		0xf5, 0x59, 0x0e, 0x41, 0x85, 0x70, 0x4d, 0x6f, 0x89, 0xcb, 0xe6, 0x36,
		0x35, 0x0d, 0x7c, 0x77, 0xe1, 0x17, 0x17, 0x5d, 0xfe, 0x48, 0x1b, 0xd6,
		0xa8, 0x86, 0xd7, 0x5b, 0x20, 0xaf, 0xb6, 0x20, 0x4b, 0x3b, 0x8e, 0xda,
		0x45, 0x4d, 0x9d, 0x84, 0xe7, 0xf2, 0x40, 0x96, 0x36, 0x89, 0xd9, 0xa1,
		0xff, 0x29, 0x71, 0x37, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x34,
		0x67, 0x0e, 0x90, 0x06, 0x00, 0x00,
	},
		"templates/index.html",
	)
}

func templates_page_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xaa, 0xae,
		0x4e, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0x55, 0x50, 0x2a, 0xc9, 0x2c, 0xc9,
		0x49, 0x55, 0xaa, 0xad, 0xe5, 0xaa, 0xae, 0xd6, 0x0b, 0x48, 0x4c, 0x4f,
		0xd5, 0x0b, 0x01, 0x09, 0xd4, 0xd6, 0x2a, 0xd4, 0x28, 0x00, 0x45, 0x9c,
		0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x61, 0x62, 0x40, 0x35, 0xa9, 0x79, 0x29,
		0x40, 0x9a, 0x0b, 0xa1, 0x3f, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x04,
		0x64, 0x82, 0x4d, 0x86, 0xa1, 0x1d, 0x9a, 0x29, 0x36, 0xfa, 0x40, 0x31,
		0xb8, 0xd1, 0x1e, 0x25, 0xb9, 0x39, 0x0a, 0x10, 0xa6, 0x33, 0x44, 0x1b,
		0x92, 0x99, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x87, 0xec, 0xa5,
		0x93, 0x00, 0x00, 0x00,
	},
		"templates/page.html",
	)
}

func templates_tag_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4c, 0x8e,
		0xc1, 0x0e, 0xc2, 0x20, 0x10, 0x44, 0xef, 0x7c, 0xc5, 0x86, 0xf4, 0x68,
		0x20, 0x5e, 0xcd, 0x96, 0x44, 0xfd, 0x04, 0x7b, 0x37, 0xa4, 0xdd, 0xb6,
		0x24, 0x48, 0x93, 0x4a, 0x0f, 0x06, 0xf9, 0x77, 0xb7, 0xad, 0x44, 0x4f,
		0xcc, 0x30, 0xf3, 0x60, 0x52, 0xea, 0xa8, 0x77, 0x81, 0x40, 0x46, 0x17,
		0x3d, 0xc9, 0x9c, 0x45, 0x4a, 0xaa, 0xb1, 0x43, 0xce, 0xf0, 0x06, 0x96,
		0xd7, 0x29, 0xf4, 0x6e, 0x50, 0xcd, 0x9a, 0x6e, 0x21, 0x85, 0x8e, 0x4f,
		0xf1, 0x03, 0xdb, 0x29, 0x44, 0x0a, 0x71, 0x45, 0x71, 0x3c, 0x9a, 0x82,
		0xa3, 0x66, 0x23, 0x70, 0xf1, 0x46, 0x00, 0x3f, 0x34, 0xdb, 0x30, 0x10,
		0x54, 0xf7, 0x03, 0x54, 0x76, 0x8e, 0xae, 0xf5, 0x04, 0xa7, 0x1a, 0xd4,
		0x79, 0xd7, 0xcf, 0xcb, 0x8b, 0x21, 0xd8, 0x49, 0xee, 0xa3, 0x77, 0x06,
		0x2d, 0x8c, 0x33, 0xf5, 0xb5, 0xd4, 0x29, 0x15, 0x46, 0xdd, 0xfc, 0xc2,
		0x0d, 0x35, 0xc6, 0x87, 0x97, 0xe6, 0xef, 0xfe, 0xbb, 0x0f, 0xb5, 0x35,
		0xa8, 0x19, 0xde, 0xfe, 0xdc, 0xa7, 0xa2, 0x5e, 0x37, 0x14, 0xf7, 0x09,
		0x00, 0x00, 0xff, 0xff, 0x72, 0xbd, 0xd1, 0x75, 0xf1, 0x00, 0x00, 0x00,
	},
		"templates/tag.html",
	)
}

func templates_base_analytics_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x92,
		0x41, 0xaf, 0xd3, 0x30, 0x0c, 0xc7, 0xef, 0x7c, 0x0a, 0xab, 0x97, 0xb4,
		0xe2, 0x29, 0xe5, 0xfc, 0xba, 0x0d, 0x6d, 0x08, 0x21, 0x2e, 0x13, 0x07,
		0x6e, 0xd3, 0x34, 0x99, 0xd4, 0xcd, 0xb2, 0x75, 0x49, 0x49, 0xdc, 0x41,
		0xd5, 0xed, 0xbb, 0xd3, 0xa6, 0x9d, 0x98, 0x10, 0x7a, 0xa7, 0xd8, 0x7f,
		0xff, 0x6a, 0xfb, 0x5f, 0xb9, 0xef, 0x4b, 0xaa, 0x8c, 0x25, 0x48, 0xd0,
		0x62, 0xdd, 0xb1, 0x51, 0x21, 0xb9, 0xdf, 0xdf, 0x01, 0xf4, 0xbd, 0xa9,
		0x40, 0x7e, 0x72, 0xb6, 0x32, 0x5a, 0x7e, 0x71, 0x4e, 0xd7, 0xb4, 0x7e,
		0x10, 0x5f, 0xcb, 0x88, 0x2c, 0x82, 0xf2, 0xa6, 0x61, 0xe0, 0xae, 0xa1,
		0x65, 0xc2, 0xf4, 0x9b, 0xf3, 0x13, 0x5e, 0x71, 0x52, 0x93, 0xd5, 0x40,
		0x00, 0x5c, 0xd1, 0xc3, 0x41, 0xe3, 0x4f, 0x58, 0x4e, 0xcf, 0xed, 0x06,
		0xbb, 0x7d, 0x11, 0x4b, 0x63, 0x2e, 0x9b, 0x36, 0x1c, 0xd3, 0x9d, 0x38,
		0x04, 0xe2, 0xb5, 0x52, 0xae, 0xb5, 0x2c, 0x5e, 0x40, 0xf4, 0xfd, 0x1b,
		0x93, 0xc5, 0x3e, 0xfb, 0x4f, 0x03, 0xf6, 0xa8, 0xce, 0xdf, 0x50, 0xd3,
		0xd5, 0xd0, 0xaf, 0x88, 0x44, 0x26, 0xad, 0x5a, 0xab, 0xd8, 0x38, 0x9b,
		0x66, 0xd0, 0x47, 0x65, 0xda, 0x49, 0xe3, 0xb0, 0x51, 0xe9, 0x54, 0x7b,
		0x21, 0xcb, 0x52, 0x79, 0x42, 0xa6, 0xcf, 0x35, 0x8d, 0x59, 0x2a, 0x26,
		0x07, 0x22, 0x2b, 0x06, 0x4c, 0x8e, 0xee, 0x06, 0x56, 0xfc, 0xe3, 0x4f,
		0xc4, 0x22, 0x86, 0xce, 0xaa, 0xa1, 0xca, 0xbe, 0xa5, 0x62, 0x6e, 0x3f,
		0xc8, 0xc1, 0x8f, 0x62, 0x2a, 0x8e, 0xcc, 0x4d, 0x78, 0x15, 0xb0, 0x7c,
		0x9a, 0x55, 0x3b, 0x85, 0xe3, 0x42, 0xb2, 0xf1, 0x8e, 0x9d, 0x72, 0x35,
		0x7c, 0x84, 0x19, 0xcc, 0x73, 0x01, 0xaf, 0x53, 0x32, 0xc6, 0x19, 0xbc,
		0x07, 0x11, 0x18, 0x39, 0x48, 0x2d, 0x4b, 0xd7, 0xfe, 0xa8, 0x49, 0xd5,
		0x46, 0x9d, 0xa5, 0x25, 0xce, 0x4b, 0x25, 0x4f, 0x41, 0x14, 0x4f, 0x96,
		0xc2, 0xb3, 0x23, 0x4d, 0x3c, 0xdb, 0x09, 0x9b, 0xee, 0x3b, 0xea, 0x2d,
		0x5e, 0xe8, 0xaf, 0xb1, 0xdd, 0x87, 0x7d, 0x01, 0x41, 0x36, 0xe8, 0x07,
		0x60, 0xeb, 0x4a, 0x92, 0xc6, 0x06, 0xf2, 0xbc, 0xa1, 0xca, 0x79, 0x4a,
		0x35, 0xbe, 0x40, 0x98, 0x7f, 0xf2, 0x3d, 0x4b, 0x63, 0xb4, 0xc8, 0xa7,
		0x8f, 0x57, 0xf1, 0x36, 0xc8, 0x8e, 0x27, 0xf0, 0x78, 0xff, 0x04, 0x00,
		0x00, 0xff, 0xff, 0xe6, 0x36, 0x02, 0xde, 0x45, 0x02, 0x00, 0x00,
	},
		"templates/base/analytics.html",
	)
}

func templates_base_base_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xd4, 0x57,
		0xcd, 0x8e, 0xdb, 0x36, 0x10, 0xbe, 0xe7, 0x29, 0x58, 0x76, 0x0f, 0x0d,
		0x1a, 0x49, 0xd9, 0x9e, 0x82, 0x42, 0x32, 0xb0, 0x75, 0x13, 0xa0, 0x87,
		0x22, 0x01, 0xd6, 0x05, 0xda, 0x53, 0x31, 0x96, 0x68, 0x89, 0x28, 0x45,
		0xaa, 0x14, 0xe5, 0xad, 0x21, 0xf8, 0xdd, 0x3b, 0xa4, 0x48, 0xfd, 0x59,
		0x59, 0x6c, 0x0e, 0x29, 0xd0, 0x93, 0x48, 0xce, 0x37, 0xff, 0x1f, 0xe9,
		0x71, 0xdf, 0x17, 0xec, 0xc4, 0x25, 0x23, 0xf4, 0x08, 0x2d, 0xa3, 0xd7,
		0xeb, 0xab, 0xf4, 0x9b, 0x9f, 0x3f, 0xee, 0x0f, 0x7f, 0x7c, 0x7a, 0x4f,
		0x2a, 0x53, 0x8b, 0xdd, 0xab, 0x74, 0xf8, 0x10, 0x92, 0x56, 0x0c, 0x0a,
		0xbb, 0xc0, 0x65, 0xcd, 0x0c, 0x90, 0xbc, 0x02, 0xdd, 0x32, 0x93, 0xd1,
		0xce, 0x9c, 0xa2, 0x77, 0x74, 0x2e, 0xaa, 0x8c, 0x69, 0x22, 0xf6, 0x77,
		0xc7, 0xcf, 0x19, 0xfd, 0x3d, 0xfa, 0xed, 0x21, 0xda, 0xab, 0xba, 0x01,
		0xc3, 0x8f, 0x82, 0x51, 0x92, 0x2b, 0x69, 0x98, 0x44, 0xbd, 0x5f, 0xde,
		0x67, 0xac, 0x28, 0xd9, 0x42, 0x53, 0x42, 0xcd, 0x32, 0x7a, 0xe6, 0xec,
		0xa9, 0x51, 0xda, 0xcc, 0xc0, 0x4f, 0xbc, 0x30, 0x55, 0x56, 0xb0, 0x33,
		0xcf, 0x59, 0xe4, 0x36, 0x6f, 0x08, 0x97, 0xdc, 0x70, 0x10, 0x51, 0x9b,
		0x83, 0x60, 0xd9, 0x7d, 0xfc, 0x16, 0x4d, 0x0d, 0xb6, 0x0c, 0x37, 0x82,
		0xed, 0xfa, 0xde, 0xb0, 0xba, 0x11, 0x60, 0x30, 0x3f, 0x77, 0x42, 0x49,
		0x7c, 0xbd, 0xa6, 0xc9, 0x20, 0xf5, 0x50, 0xc1, 0xe5, 0x5f, 0xa4, 0xd2,
		0xec, 0x94, 0xd1, 0xe4, 0xc4, 0x58, 0xd1, 0x26, 0x20, 0x44, 0x0c, 0x46,
		0xd5, 0xf1, 0x3f, 0xb5, 0xa0, 0xc4, 0x5c, 0x1a, 0x8c, 0x08, 0x9a, 0x46,
		0xf0, 0x1c, 0x53, 0x50, 0x32, 0xb1, 0xb2, 0xef, 0x9d, 0x4c, 0x33, 0x81,
		0x22, 0x61, 0x98, 0x96, 0xe8, 0x04, 0xb1, 0xd6, 0x70, 0x46, 0x1f, 0xb9,
		0x61, 0x18, 0x22, 0x23, 0x0f, 0x87, 0x8f, 0xbf, 0x92, 0x0f, 0x68, 0xd4,
		0xe7, 0xd8, 0xf7, 0xfc, 0x44, 0xe2, 0xbd, 0x92, 0x27, 0x5e, 0xc6, 0x1f,
		0x00, 0x73, 0x51, 0x12, 0x6b, 0x3e, 0xc5, 0xe1, 0x0c, 0xb6, 0x15, 0xa6,
		0x9e, 0x77, 0x86, 0x58, 0x31, 0xf5, 0xb1, 0xf5, 0xfd, 0x8d, 0x1e, 0x25,
		0x49, 0x30, 0xcb, 0x64, 0x81, 0x76, 0xfc, 0x66, 0x4a, 0xda, 0xb6, 0x8c,
		0x69, 0x97, 0xb5, 0xed, 0x60, 0x32, 0xb4, 0xd0, 0x2e, 0x8f, 0xaa, 0xb8,
		0xf8, 0xba, 0x17, 0xfc, 0x4c, 0x72, 0x01, 0x6d, 0x9b, 0x51, 0x09, 0xe7,
		0x23, 0x68, 0x32, 0x7c, 0x22, 0xa4, 0x06, 0x74, 0xc2, 0x84, 0x6d, 0x6b,
		0x30, 0xfd, 0x3c, 0x32, 0xaa, 0xc1, 0xc4, 0x95, 0xcd, 0x13, 0x05, 0xbc,
		0x74, 0x35, 0xf1, 0xf9, 0x2d, 0xad, 0xd9, 0xde, 0x01, 0x92, 0x4b, 0x8f,
		0xd2, 0x2d, 0x6f, 0x91, 0x0f, 0x72, 0xc2, 0xd8, 0xf0, 0x3a, 0x63, 0x94,
		0xf4, 0xb5, 0x1f, 0x36, 0x74, 0xa5, 0x66, 0x54, 0x59, 0xda, 0x8e, 0x16,
		0x60, 0xc0, 0x6f, 0xac, 0x4b, 0x21, 0xa0, 0x69, 0xc7, 0x63, 0xd0, 0xa5,
		0x65, 0x68, 0xec, 0x75, 0x46, 0xf1, 0xdc, 0x19, 0xba, 0x6b, 0x1b, 0x90,
		0xc1, 0x7c, 0xab, 0x23, 0x25, 0xc5, 0x85, 0xee, 0x0e, 0xce, 0x26, 0x99,
		0x92, 0x4c, 0x13, 0x8b, 0x7b, 0x46, 0xd5, 0xb6, 0x25, 0x42, 0x3f, 0x74,
		0xf7, 0x5f, 0x41, 0xd3, 0x64, 0x28, 0xce, 0xe2, 0x0c, 0x56, 0x95, 0x3a,
		0x6a, 0x90, 0x45, 0xe0, 0x51, 0x42, 0x77, 0x13, 0x95, 0x0e, 0x96, 0xaf,
		0xf6, 0x46, 0xc0, 0xac, 0x43, 0x09, 0xb6, 0x68, 0xda, 0x3a, 0xca, 0x2a,
		0x4d, 0xe2, 0x4f, 0x50, 0xb2, 0x96, 0x7c, 0x87, 0xa6, 0x48, 0x7c, 0x80,
		0xb2, 0x1d, 0x79, 0xfc, 0x58, 0xa9, 0x27, 0x7b, 0xf0, 0xda, 0x0b, 0xf7,
		0x48, 0xbd, 0x52, 0x69, 0xce, 0x96, 0x90, 0xe9, 0x38, 0x00, 0x1f, 0x34,
		0xf2, 0x49, 0xac, 0x60, 0x0f, 0x3a, 0xaf, 0xf8, 0x99, 0xbd, 0xf6, 0xd7,
		0xe2, 0x33, 0x9c, 0x09, 0x8d, 0x24, 0x9b, 0x1d, 0xdd, 0x8c, 0xf9, 0xa5,
		0xde, 0xd0, 0x5f, 0x27, 0x66, 0xee, 0x02, 0xfd, 0xf1, 0xb3, 0xa2, 0xcd,
		0xe2, 0x36, 0xcf, 0xac, 0x2d, 0x8c, 0xb9, 0x9b, 0xbd, 0xc3, 0x9e, 0xf8,
		0xf2, 0xc3, 0x80, 0x69, 0x63, 0xfb, 0xae, 0xd2, 0x9d, 0x57, 0x69, 0x6d,
		0x0b, 0xd2, 0x04, 0x91, 0x2b, 0x0f, 0xc3, 0xc5, 0x5e, 0x9e, 0x61, 0x3b,
		0x4b, 0x46, 0xee, 0xfe, 0x7c, 0x43, 0xee, 0x1a, 0x4c, 0x90, 0xfc, 0x98,
		0xf9, 0x4c, 0x9f, 0x75, 0x6c, 0xa1, 0x6d, 0xd2, 0xf7, 0x4e, 0x27, 0x7e,
		0x14, 0x5d, 0x79, 0xbd, 0xfa, 0x20, 0xc2, 0xe1, 0x8c, 0x0e, 0x2f, 0x89,
		0x25, 0x4d, 0x3a, 0xb1, 0x2c, 0x3b, 0x41, 0x08, 0x59, 0x60, 0x42, 0x2b,
		0xbe, 0x90, 0x1a, 0xdb, 0x04, 0x7b, 0x59, 0x97, 0xc2, 0x52, 0xf3, 0xb2,
		0x32, 0x5b, 0x2d, 0x7b, 0x59, 0x28, 0xb7, 0xc5, 0x0c, 0xee, 0x0a, 0xad,
		0x9a, 0x42, 0x3d, 0xc9, 0xf5, 0x2b, 0x12, 0x4a, 0xfd, 0xed, 0xea, 0x51,
		0x9a, 0xf0, 0x33, 0xa7, 0xe9, 0x71, 0x7c, 0x26, 0x41, 0x33, 0x63, 0x2f,
		0xf7, 0x71, 0xb7, 0xb8, 0x88, 0xab, 0x24, 0x83, 0x95, 0xa8, 0x66, 0xb2,
		0x5b, 0xb9, 0x5e, 0xb1, 0x22, 0x1f, 0xdc, 0x5c, 0x1c, 0x33, 0x3e, 0x9b,
		0x91, 0x55, 0xba, 0x83, 0x70, 0x2b, 0x10, 0x7a, 0x37, 0xde, 0x91, 0x9f,
		0x2e, 0xfb, 0x60, 0x62, 0x34, 0x76, 0xa3, 0x9d, 0xae, 0x19, 0x32, 0x2b,
		0x41, 0x12, 0xb4, 0x2c, 0xe1, 0x26, 0x0b, 0x13, 0xdf, 0xa6, 0xb3, 0xe5,
		0x63, 0x77, 0x04, 0x9c, 0x08, 0x48, 0xd3, 0x09, 0x11, 0xfa, 0xd7, 0xf7,
		0x82, 0x49, 0x32, 0x06, 0x6a, 0xe9, 0xe9, 0x5e, 0xc1, 0x9b, 0x5a, 0x59,
		0x42, 0xde, 0x86, 0xb4, 0x75, 0x85, 0xd6, 0xc4, 0xdd, 0xd2, 0xdc, 0xbe,
		0x7a, 0x81, 0x3d, 0x9b, 0xfc, 0xfc, 0x1a, 0x8c, 0x71, 0x8e, 0xbe, 0x26,
		0x57, 0x0c, 0x94, 0x8e, 0x26, 0x1b, 0x09, 0x3c, 0x4b, 0x10, 0xc4, 0x3b,
		0xe5, 0x2f, 0xa3, 0x05, 0x2a, 0x58, 0x46, 0x38, 0xbd, 0x89, 0x0c, 0x6e,
		0xfb, 0xbf, 0xe1, 0xc1, 0xed, 0xb3, 0xb7, 0x44, 0xdc, 0xfc, 0x80, 0xce,
		0xc5, 0x33, 0xa1, 0x5f, 0xde, 0x0c, 0x61, 0xe3, 0xd8, 0x64, 0x7f, 0xdc,
		0xa2, 0xba, 0x88, 0xee, 0xdf, 0x86, 0x95, 0x3a, 0x9d, 0x70, 0xe2, 0x8e,
		0xee, 0xdd, 0x5e, 0x94, 0xd1, 0xbb, 0xb0, 0xf0, 0x82, 0x1f, 0xc6, 0x5e,
		0xcf, 0xa7, 0x40, 0x3f, 0x43, 0x87, 0x31, 0x70, 0x11, 0xc4, 0x1c, 0x77,
		0x52, 0xca, 0x4c, 0xd3, 0xe2, 0x52, 0x06, 0x12, 0xc4, 0x05, 0x2b, 0xdf,
		0x4e, 0xc3, 0xe4, 0x30, 0x42, 0xe2, 0x50, 0xe9, 0xfe, 0x20, 0x84, 0x34,
		0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xef, 0xd0, 0xa6, 0xc6, 0x52, 0x0c,
		0x00, 0x00,
	},
		"templates/base/base.html",
	)
}

func templates_base_footer_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x92,
		0xdd, 0x8e, 0xa3, 0x30, 0x0c, 0x85, 0xef, 0xf7, 0x29, 0xac, 0x5c, 0x56,
		0x82, 0x94, 0xbd, 0x5a, 0xad, 0x28, 0xda, 0x87, 0xd8, 0x17, 0x08, 0x60,
		0x20, 0x28, 0x24, 0xd9, 0xd8, 0xdd, 0xfe, 0xa0, 0xbe, 0xfb, 0x18, 0x5a,
		0xa6, 0x33, 0x53, 0xcd, 0x0d, 0x3a, 0x8e, 0xc3, 0xe7, 0x73, 0xc0, 0xf3,
		0xdc, 0x62, 0x67, 0x3d, 0x82, 0xea, 0x42, 0x60, 0x4c, 0xea, 0x76, 0xfb,
		0x01, 0x30, 0xcf, 0x7a, 0x07, 0x7f, 0x07, 0x4b, 0x40, 0x43, 0x38, 0xba,
		0x16, 0x6a, 0x04, 0x8a, 0xce, 0x32, 0x04, 0x0f, 0x1e, 0x4f, 0x90, 0xd0,
		0xa1, 0x21, 0x24, 0xd8, 0xe9, 0xf5, 0x85, 0xb2, 0xb5, 0xff, 0xa1, 0x71,
		0x86, 0xe8, 0xa0, 0x9a, 0xe0, 0xd9, 0x08, 0x32, 0x41, 0x13, 0x5c, 0x36,
		0xb5, 0x59, 0xb1, 0xdf, 0x54, 0xe8, 0x3a, 0x42, 0xce, 0x8a, 0xb5, 0x76,
		0x7d, 0xf6, 0x6b, 0x13, 0x8f, 0xc6, 0x4f, 0x55, 0x09, 0xec, 0x33, 0x8e,
		0xf1, 0xcc, 0x59, 0x83, 0x7e, 0x31, 0x77, 0xef, 0x4a, 0x3f, 0x6e, 0x4a,
		0xb4, 0x81, 0x21, 0x61, 0x77, 0x50, 0x03, 0x73, 0xfc, 0xad, 0x75, 0x6f,
		0x79, 0x38, 0xd6, 0x79, 0x13, 0x26, 0x6d, 0xfa, 0xe0, 0xaf, 0xc6, 0xe1,
		0x35, 0x05, 0x1d, 0x83, 0x0b, 0xaa, 0x5a, 0x9e, 0xa5, 0x36, 0x15, 0x48,
		0xb4, 0xc9, 0xb4, 0x08, 0x27, 0xb9, 0xfd, 0x44, 0x51, 0x34, 0x7e, 0x9b,
		0xdb, 0xbb, 0x4b, 0x1c, 0xac, 0x84, 0x81, 0x77, 0x95, 0x0d, 0x68, 0x12,
		0xab, 0xaa, 0xd4, 0xcb, 0xc5, 0x0a, 0xea, 0xcb, 0xb7, 0x2e, 0x58, 0xc0,
		0xe2, 0xf8, 0xab, 0x0d, 0x55, 0xfd, 0xf9, 0x50, 0x2d, 0x4e, 0xb6, 0x44,
		0xfa, 0x11, 0xa9, 0xd4, 0x92, 0x7d, 0x91, 0x4f, 0x41, 0x4d, 0xb2, 0x91,
		0x81, 0x52, 0x73, 0xc7, 0x93, 0xf0, 0xcd, 0x68, 0xce, 0x79, 0x1f, 0x42,
		0x2f, 0xff, 0x21, 0x5a, 0xba, 0xcf, 0x91, 0x33, 0xed, 0x6c, 0x4d, 0x7a,
		0xfc, 0x77, 0xc4, 0x74, 0xd1, 0x45, 0x5e, 0x14, 0xf9, 0xfe, 0x51, 0xe5,
		0x93, 0xf5, 0xf9, 0x48, 0xab, 0xfb, 0x15, 0xf8, 0xc2, 0x96, 0x6f, 0x87,
		0x5c, 0xcb, 0x1a, 0x10, 0x27, 0x13, 0x57, 0x64, 0x6b, 0x89, 0xf5, 0x48,
		0xfa, 0x79, 0xfa, 0x4a, 0x99, 0x67, 0xf4, 0xad, 0x6c, 0xc1, 0x5b, 0x00,
		0x00, 0x00, 0xff, 0xff, 0xa9, 0xef, 0x18, 0x09, 0x4b, 0x02, 0x00, 0x00,
	},
		"templates/base/footer.html",
	)
}

func templates_base_header_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0xcf,
		0x31, 0x0e, 0x02, 0x21, 0x10, 0x85, 0xe1, 0xde, 0x53, 0x90, 0xe9, 0x05,
		0x93, 0x6d, 0x5d, 0xef, 0x82, 0xcc, 0xdb, 0x40, 0x84, 0xc1, 0x30, 0xd3,
		0x18, 0xb2, 0x77, 0x77, 0x2b, 0xf5, 0x02, 0xb6, 0x7f, 0x5e, 0xf2, 0xe5,
		0xcd, 0xc9, 0xd8, 0x8a, 0xc0, 0x51, 0x46, 0x64, 0x0c, 0xda, 0xf7, 0x93,
		0x73, 0xd7, 0x5a, 0xe4, 0xe1, 0x06, 0xea, 0x4a, 0x6a, 0xaf, 0x0a, 0xcd,
		0x80, 0x91, 0xcb, 0x03, 0xdb, 0x4a, 0x21, 0x08, 0x8c, 0x25, 0xfa, 0x7b,
		0xef, 0xa6, 0x36, 0xe2, 0x33, 0xb1, 0xf8, 0xd4, 0x5b, 0xf8, 0x84, 0xb0,
		0xf8, 0x8b, 0x5f, 0x42, 0x52, 0xfd, 0x36, 0xdf, 0xca, 0xb1, 0x52, 0xa5,
		0xdb, 0x9f, 0x80, 0xb3, 0x65, 0x34, 0xfc, 0x30, 0x73, 0x42, 0xf8, 0xf8,
		0xf3, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x24, 0xeb, 0x71, 0xe4, 0x00,
		0x00, 0x00,
	},
		"templates/base/header.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/archives.html": templates_archives_html,
	"templates/article.html": templates_article_html,
	"templates/atom.xml": templates_atom_xml,
	"templates/category.html": templates_category_html,
	"templates/index.html": templates_index_html,
	"templates/page.html": templates_page_html,
	"templates/tag.html": templates_tag_html,
	"templates/base/analytics.html": templates_base_analytics_html,
	"templates/base/base.html": templates_base_base_html,
	"templates/base/footer.html": templates_base_footer_html,
	"templates/base/header.html": templates_base_header_html,
}