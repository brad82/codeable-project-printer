// Ported heavily from https://github.com/song940/node-escpos/blob/v3/packages/printer/src/commands.ts
package escpos

const LF = "\x0a"
const FS = "\x1c"
const FF = "\x0c"
const GS = "\x1d"
const DLE = "\x10"
const EOT = "\x04"
const NUL = "\x00"
const ESC = "\x1b"
const TAB = "\x74"
const EOL = "\n"

const (
	CTL_LF  = "\x0a"     // Print and line feed
	CTL_GLF = "\x4a\x00" // Print and feed paper (without spaces between lines)
	CTL_FF  = "\x0c"     // Form feed
	CTL_CR  = "\x0d"     // Carriage return
	CTL_HT  = "\x09"     // Horizontal tab
	CTL_VT  = "\x0b"     // Vertical tab
)

const (
	CS_DEFAULT = "\x1b\x20\x00"
	CS_SET     = "\x1b\x20"
)

const (
	LS_DEFAULT = "\x1b\x32"
	LS_SET     = "\x1b\x33"
)

const (
	HW_INIT   = "\x1b\x40"         // Clear data in buffer and reset modes
	HW_SELECT = "\x1b\x3d\x01"     // Printer select
	HW_RESET  = "\x1b\x3f\x0a\x00" // Reset printer hardware
)

const (
	CD_KICK_2 = "\x1b\x70\x00\x19\x78" // Sends a pulse to pin 2 []
	CD_KICK_5 = "\x1b\x70\x01\x19\x78" // Sends a pulse to pin 5 []
)

const (
	BOTTOM = "\x1b\x4f" // Fix bottom size
	LEFT   = "\x1b\x6c" // Fix left size
	RIGHT  = "\x1b\x51" // Fix right size
)

const (
	PAPER_FULL_CUT = "\x1d\x56\x00" // Full cut paper
	PAPER_PART_CUT = "\x1d\x56\x01" // Partial cut paper
	PAPER_CUT_A    = "\x1d\x56\x41" // Partial cut paper
	PAPER_CUT_B    = "\x1d\x56\x42" // Partial cut paper
	STAR_FULL_CUT  = "\x1B\x64\x02" // STAR printer - Full cut
)

var TXT_HEIGHT = map[int]string{
	1: "\x00",
	2: "\x01",
	3: "\x02",
	4: "\x03",
	5: "\x04",
	6: "\x05",
	7: "\x06",
	8: "\x07",
}

var TXT_WIDTH = map[int]string{
	1: "\x00",
	2: "\x10",
	3: "\x20",
	4: "\x30",
	5: "\x40",
	6: "\x50",
	7: "\x60",
	8: "\x70",
}

const (
	TXT_NORMAL                 = "\x1b\x21\x00" // Normal text
	TXT_2HEIGHT                = "\x1b\x21\x10" // Double height text
	TXT_2WIDTH                 = "\x1b\x21\x20" // Double width text
	TXT_4SQUARE                = "\x1b\x21\x30" // Double width & height text
	STAR_TXT_EMPHASIZED        = "\x1B\x45"     // STAR printer - Select emphasized printing
	STAR_CANCEL_TXT_EMPHASIZED = "\x1B\x46"     // STAR printer - Cancel emphasized printing

	TXT_UNDERL_OFF = "\x1b\x2d\x00" // Underline font OFF
	TXT_UNDERL_ON  = "\x1b\x2d\x01" // Underline font 1-dot ON
	TXT_UNDERL2_ON = "\x1b\x2d\x02" // Underline font 2-dot ON
	TXT_BOLD_OFF   = "\x1b\x45\x00" // Bold font OFF
	TXT_BOLD_ON    = "\x1b\x45\x01" // Bold font ON
	TXT_ITALIC_OFF = "\x1b\x35"     // Italic font ON
	TXT_ITALIC_ON  = "\x1b\x34"     // Italic font ON

	TXT_FONT_A = "\x1b\x4d\x00" // Font type A
	TXT_FONT_B = "\x1b\x4d\x01" // Font type B
	TXT_FONT_C = "\x1b\x4d\x02" // Font type C

	TXT_ALIGN_LT = "\x1b\x61\x00" // Left justification
	TXT_ALIGN_CT = "\x1b\x61\x01" // Centering
	TXT_ALIGN_RT = "\x1b\x61\x02" // Right justification

	STAR_TXT_ALIGN_LA = "\x1B\x1D\x61\x00" // STAR printer - Left alignment
	STAR_TXT_ALIGN_CA = "\x1B\x1D\x61\x01" // STAR printer - Center alignment
	STAR_TXT_ALIGN_RA = "\x1B\x1D\x61\x02" // STAR printer - Right alignment
)

var BARCODE_WIDTH = map[int]string{
	1: "\x1d\x77\x02",
	2: "\x1d\x77\x03",
	3: "\x1d\x77\x04",
	4: "\x1d\x77\x05",
	5: "\x1d\x77\x06",
}

const (
	BARCODE_TXT_OFF = "\x1d\x48\x00" // HRI barcode chars OFF
	BARCODE_TXT_ABV = "\x1d\x48\x01" // HRI barcode chars above
	BARCODE_TXT_BLW = "\x1d\x48\x02" // HRI barcode chars below
	BARCODE_TXT_BTH = "\x1d\x48\x03" // HRI barcode chars both above and below

	BARCODE_FONT_A = "\x1d\x66\x00" // Font type A for HRI barcode chars
	BARCODE_FONT_B = "\x1d\x66\x01" // Font type B for HRI barcode chars

	/*
		 BARCODE_HEIGHT: function (height: number) { // Barcode Height [1-255]
			 return Buffer.from("1d68"+ numToHexString(height), "hex");
		 },
	*/

	// Barcode Width  [2-6]

	BARCODE_HEIGHT_DEFAULT = "\x1d\x68\x64" // Barcode height default:100
	BARCODE_WIDTH_DEFAULT  = "\x1d\x77\x01" // Barcode width default:1

	BARCODE_UPC_A   = "\x1d\x6b\x00" // Barcode type UPC-A
	BARCODE_UPC_E   = "\x1d\x6b\x01" // Barcode type UPC-E
	BARCODE_EAN13   = "\x1d\x6b\x02" // Barcode type EAN13
	BARCODE_EAN8    = "\x1d\x6b\x03" // Barcode type EAN8
	BARCODE_CODE39  = "\x1d\x6b\x04" // Barcode type CODE39
	BARCODE_ITF     = "\x1d\x6b\x05" // Barcode type ITF
	BARCODE_NW7     = "\x1d\x6b\x06" // Barcode type NW7
	BARCODE_CODE93  = "\x1d\x6b\x48" // Barcode type CODE93
	BARCODE_CODE128 = "\x1d\x6b\x49" // Barcode type CODE128
)

const (
	CODE2D_TYPE_PDF417     = GS + "Z" + "\x00"
	CODE2D_TYPE_DATAMATRIX = GS + "Z" + "\x01"
	CODE2D_TYPE_QR         = GS + "Z" + "\x02"
	CODE2D_CODE2D          = ESC + "Z"
	CODE2D_QR_LEVEL_L      = "L" // correct level 7%
	CODE2D_QR_LEVEL_M      = "M" // correct level 15%
	CODE2D_QR_LEVEL_Q      = "Q" // correct level 25%
	CODE2D_QR_LEVEL_H      = "H" // correct level 30%
)

const (
	S_RASTER_N  = "\x1d\x76\x30\x00" // Set raster image normal size
	S_RASTER_2W = "\x1d\x76\x30\x01" // Set raster image double width
	S_RASTER_2H = "\x1d\x76\x30\x02" // Set raster image double height
	S_RASTER_Q  = "\x1d\x76\x30\x03" // Set raster image quadruple
)

const (
	BITMAP_S8  = "\x1b\x2a\x00"
	BITMAP_D8  = "\x1b\x2a\x01"
	BITMAP_S24 = "\x1b\x2a\x20"
	BITMAP_D24 = "\x1b\x2a\x21"
)

const (
	GSV0_NORMAL = "\x1d\x76\x30\x00"
	GSV0_DW     = "\x1d\x76\x30\x01"
	GSV0_DH     = "\x1d\x76\x30\x02"
	GSV0_DWDH   = "\x1d\x76\x30\x03"
)

const BEEP = "\x1b\x42" // Printer Buzzer pre hex

const (
	COLOR_BLACK     = "\x1b\x72\x00" // black
	COLOR_RED       = "\x1b\x72\x01" // red
	COLOR_REVERSE   = "\x1dB1"       // Reverses the colors - white text on black background
	COLOR_UNREVERSE = "\x1dB0"       // Default: undo the reverse - black text on white background
)
