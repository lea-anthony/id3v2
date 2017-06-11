// Code generated by `go run generate_ids.go`. DO NOT EDIT.

// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a Modified
// BSD License that can be found in the LICENSE file.

package id3v2

// These are the standard frame ids as specified in the
// v2.4.0 and v2.3.0 specifications.
const (
	FrameAENC FrameID = 'A'<<24 | 'E'<<16 | 'N'<<8 | 'C' // Audio encryption
	FrameAPIC FrameID = 'A'<<24 | 'P'<<16 | 'I'<<8 | 'C' // Attached picture
	FrameASPI FrameID = 'A'<<24 | 'S'<<16 | 'P'<<8 | 'I' // Audio seek point index
	FrameCOMM FrameID = 'C'<<24 | 'O'<<16 | 'M'<<8 | 'M' // Comments
	FrameCOMR FrameID = 'C'<<24 | 'O'<<16 | 'M'<<8 | 'R' // Commercial frame
	FrameENCR FrameID = 'E'<<24 | 'N'<<16 | 'C'<<8 | 'R' // Encryption method registration
	FrameEQU2 FrameID = 'E'<<24 | 'Q'<<16 | 'U'<<8 | '2' // Equalisation (2)
	FrameETCO FrameID = 'E'<<24 | 'T'<<16 | 'C'<<8 | 'O' // Event timing codes
	FrameGEOB FrameID = 'G'<<24 | 'E'<<16 | 'O'<<8 | 'B' // General encapsulated object
	FrameGRID FrameID = 'G'<<24 | 'R'<<16 | 'I'<<8 | 'D' // Group identification registration
	FrameLINK FrameID = 'L'<<24 | 'I'<<16 | 'N'<<8 | 'K' // Linked information
	FrameMCDI FrameID = 'M'<<24 | 'C'<<16 | 'D'<<8 | 'I' // Music CD identifier
	FrameMLLT FrameID = 'M'<<24 | 'L'<<16 | 'L'<<8 | 'T' // MPEG location lookup table
	FrameOWNE FrameID = 'O'<<24 | 'W'<<16 | 'N'<<8 | 'E' // Ownership frame
	FramePRIV FrameID = 'P'<<24 | 'R'<<16 | 'I'<<8 | 'V' // Private frame
	FramePCNT FrameID = 'P'<<24 | 'C'<<16 | 'N'<<8 | 'T' // Play counter
	FramePOPM FrameID = 'P'<<24 | 'O'<<16 | 'P'<<8 | 'M' // Popularimeter
	FramePOSS FrameID = 'P'<<24 | 'O'<<16 | 'S'<<8 | 'S' // Position synchronisation frame
	FrameRBUF FrameID = 'R'<<24 | 'B'<<16 | 'U'<<8 | 'F' // Recommended buffer size
	FrameRVA2 FrameID = 'R'<<24 | 'V'<<16 | 'A'<<8 | '2' // Relative volume adjustment (2)
	FrameRVRB FrameID = 'R'<<24 | 'V'<<16 | 'R'<<8 | 'B' // Reverb
	FrameSEEK FrameID = 'S'<<24 | 'E'<<16 | 'E'<<8 | 'K' // Seek frame
	FrameSIGN FrameID = 'S'<<24 | 'I'<<16 | 'G'<<8 | 'N' // Signature frame
	FrameSYLT FrameID = 'S'<<24 | 'Y'<<16 | 'L'<<8 | 'T' // Synchronised lyric/text
	FrameSYTC FrameID = 'S'<<24 | 'Y'<<16 | 'T'<<8 | 'C' // Synchronised tempo codes
	FrameTALB FrameID = 'T'<<24 | 'A'<<16 | 'L'<<8 | 'B' // Album/Movie/Show title
	FrameTBPM FrameID = 'T'<<24 | 'B'<<16 | 'P'<<8 | 'M' // BPM (beats per minute)
	FrameTCOM FrameID = 'T'<<24 | 'C'<<16 | 'O'<<8 | 'M' // Composer
	FrameTCON FrameID = 'T'<<24 | 'C'<<16 | 'O'<<8 | 'N' // Content type
	FrameTCOP FrameID = 'T'<<24 | 'C'<<16 | 'O'<<8 | 'P' // Copyright message
	FrameTDEN FrameID = 'T'<<24 | 'D'<<16 | 'E'<<8 | 'N' // Encoding time
	FrameTDLY FrameID = 'T'<<24 | 'D'<<16 | 'L'<<8 | 'Y' // Playlist delay
	FrameTDOR FrameID = 'T'<<24 | 'D'<<16 | 'O'<<8 | 'R' // Original release time
	FrameTDRC FrameID = 'T'<<24 | 'D'<<16 | 'R'<<8 | 'C' // Recording time
	FrameTDRL FrameID = 'T'<<24 | 'D'<<16 | 'R'<<8 | 'L' // Release time
	FrameTDTG FrameID = 'T'<<24 | 'D'<<16 | 'T'<<8 | 'G' // Tagging time
	FrameTENC FrameID = 'T'<<24 | 'E'<<16 | 'N'<<8 | 'C' // Encoded by
	FrameTEXT FrameID = 'T'<<24 | 'E'<<16 | 'X'<<8 | 'T' // Lyricist/Text writer
	FrameTFLT FrameID = 'T'<<24 | 'F'<<16 | 'L'<<8 | 'T' // File type
	FrameTIPL FrameID = 'T'<<24 | 'I'<<16 | 'P'<<8 | 'L' // Involved people list
	FrameTIT1 FrameID = 'T'<<24 | 'I'<<16 | 'T'<<8 | '1' // Content group description
	FrameTIT2 FrameID = 'T'<<24 | 'I'<<16 | 'T'<<8 | '2' // Title/songname/content description
	FrameTIT3 FrameID = 'T'<<24 | 'I'<<16 | 'T'<<8 | '3' // Subtitle/Description refinement
	FrameTKEY FrameID = 'T'<<24 | 'K'<<16 | 'E'<<8 | 'Y' // Initial key
	FrameTLAN FrameID = 'T'<<24 | 'L'<<16 | 'A'<<8 | 'N' // Language(s)
	FrameTLEN FrameID = 'T'<<24 | 'L'<<16 | 'E'<<8 | 'N' // Length
	FrameTMCL FrameID = 'T'<<24 | 'M'<<16 | 'C'<<8 | 'L' // Musician credits list
	FrameTMED FrameID = 'T'<<24 | 'M'<<16 | 'E'<<8 | 'D' // Media type
	FrameTMOO FrameID = 'T'<<24 | 'M'<<16 | 'O'<<8 | 'O' // Mood
	FrameTOAL FrameID = 'T'<<24 | 'O'<<16 | 'A'<<8 | 'L' // Original album/movie/show title
	FrameTOFN FrameID = 'T'<<24 | 'O'<<16 | 'F'<<8 | 'N' // Original filename
	FrameTOLY FrameID = 'T'<<24 | 'O'<<16 | 'L'<<8 | 'Y' // Original lyricist(s)/text writer(s)
	FrameTOPE FrameID = 'T'<<24 | 'O'<<16 | 'P'<<8 | 'E' // Original artist(s)/performer(s)
	FrameTOWN FrameID = 'T'<<24 | 'O'<<16 | 'W'<<8 | 'N' // File owner/licensee
	FrameTPE1 FrameID = 'T'<<24 | 'P'<<16 | 'E'<<8 | '1' // Lead performer(s)/Soloist(s)
	FrameTPE2 FrameID = 'T'<<24 | 'P'<<16 | 'E'<<8 | '2' // Band/orchestra/accompaniment
	FrameTPE3 FrameID = 'T'<<24 | 'P'<<16 | 'E'<<8 | '3' // Conductor/performer refinement
	FrameTPE4 FrameID = 'T'<<24 | 'P'<<16 | 'E'<<8 | '4' // Interpreted, remixed, or otherwise modified by
	FrameTPOS FrameID = 'T'<<24 | 'P'<<16 | 'O'<<8 | 'S' // Part of a set
	FrameTPRO FrameID = 'T'<<24 | 'P'<<16 | 'R'<<8 | 'O' // Produced notice
	FrameTPUB FrameID = 'T'<<24 | 'P'<<16 | 'U'<<8 | 'B' // Publisher
	FrameTRCK FrameID = 'T'<<24 | 'R'<<16 | 'C'<<8 | 'K' // Track number/Position in set
	FrameTRSN FrameID = 'T'<<24 | 'R'<<16 | 'S'<<8 | 'N' // Internet radio station name
	FrameTRSO FrameID = 'T'<<24 | 'R'<<16 | 'S'<<8 | 'O' // Internet radio station owner
	FrameTSOA FrameID = 'T'<<24 | 'S'<<16 | 'O'<<8 | 'A' // Album sort order
	FrameTSOP FrameID = 'T'<<24 | 'S'<<16 | 'O'<<8 | 'P' // Performer sort order
	FrameTSOT FrameID = 'T'<<24 | 'S'<<16 | 'O'<<8 | 'T' // Title sort order
	FrameTSRC FrameID = 'T'<<24 | 'S'<<16 | 'R'<<8 | 'C' // ISRC (international standard recording code)
	FrameTSSE FrameID = 'T'<<24 | 'S'<<16 | 'S'<<8 | 'E' // Software/Hardware and settings used for encoding
	FrameTSST FrameID = 'T'<<24 | 'S'<<16 | 'S'<<8 | 'T' // Set subtitle
	FrameTXXX FrameID = 'T'<<24 | 'X'<<16 | 'X'<<8 | 'X' // User defined text information frame
	FrameUFID FrameID = 'U'<<24 | 'F'<<16 | 'I'<<8 | 'D' // Unique file identifier
	FrameUSER FrameID = 'U'<<24 | 'S'<<16 | 'E'<<8 | 'R' // Terms of use
	FrameUSLT FrameID = 'U'<<24 | 'S'<<16 | 'L'<<8 | 'T' // Unsynchronised lyric/text transcription
	FrameWCOM FrameID = 'W'<<24 | 'C'<<16 | 'O'<<8 | 'M' // Commercial information
	FrameWCOP FrameID = 'W'<<24 | 'C'<<16 | 'O'<<8 | 'P' // Copyright/Legal information
	FrameWOAF FrameID = 'W'<<24 | 'O'<<16 | 'A'<<8 | 'F' // Official audio file webpage
	FrameWOAR FrameID = 'W'<<24 | 'O'<<16 | 'A'<<8 | 'R' // Official artist/performer webpage
	FrameWOAS FrameID = 'W'<<24 | 'O'<<16 | 'A'<<8 | 'S' // Official audio source webpage
	FrameWORS FrameID = 'W'<<24 | 'O'<<16 | 'R'<<8 | 'S' // Official Internet radio station homepage
	FrameWPAY FrameID = 'W'<<24 | 'P'<<16 | 'A'<<8 | 'Y' // Payment
	FrameWPUB FrameID = 'W'<<24 | 'P'<<16 | 'U'<<8 | 'B' // Publishers official webpage
	FrameWXXX FrameID = 'W'<<24 | 'X'<<16 | 'X'<<8 | 'X' // User defined URL link frame
	FrameEQUA FrameID = 'E'<<24 | 'Q'<<16 | 'U'<<8 | 'A' // Equalization
	FrameIPLS FrameID = 'I'<<24 | 'P'<<16 | 'L'<<8 | 'S' // Involved people list
	FrameRVAD FrameID = 'R'<<24 | 'V'<<16 | 'A'<<8 | 'D' // Relative volume adjustment
	FrameTDAT FrameID = 'T'<<24 | 'D'<<16 | 'A'<<8 | 'T' // Date
	FrameTIME FrameID = 'T'<<24 | 'I'<<16 | 'M'<<8 | 'E' // Time
	FrameTORY FrameID = 'T'<<24 | 'O'<<16 | 'R'<<8 | 'Y' // Original release year
	FrameTRDA FrameID = 'T'<<24 | 'R'<<16 | 'D'<<8 | 'A' // Recording dates
	FrameTSIZ FrameID = 'T'<<24 | 'S'<<16 | 'I'<<8 | 'Z' // Size
	FrameTYER FrameID = 'T'<<24 | 'Y'<<16 | 'E'<<8 | 'R' // Year
)

func (id FrameID) String() string {
	switch id {
	case FrameAENC:
		return "AENC: Audio encryption"
	case FrameAPIC:
		return "APIC: Attached picture"
	case FrameASPI:
		return "ASPI: Audio seek point index"
	case FrameCOMM:
		return "COMM: Comments"
	case FrameCOMR:
		return "COMR: Commercial frame"
	case FrameENCR:
		return "ENCR: Encryption method registration"
	case FrameEQU2:
		return "EQU2: Equalisation (2)"
	case FrameETCO:
		return "ETCO: Event timing codes"
	case FrameGEOB:
		return "GEOB: General encapsulated object"
	case FrameGRID:
		return "GRID: Group identification registration"
	case FrameLINK:
		return "LINK: Linked information"
	case FrameMCDI:
		return "MCDI: Music CD identifier"
	case FrameMLLT:
		return "MLLT: MPEG location lookup table"
	case FrameOWNE:
		return "OWNE: Ownership frame"
	case FramePRIV:
		return "PRIV: Private frame"
	case FramePCNT:
		return "PCNT: Play counter"
	case FramePOPM:
		return "POPM: Popularimeter"
	case FramePOSS:
		return "POSS: Position synchronisation frame"
	case FrameRBUF:
		return "RBUF: Recommended buffer size"
	case FrameRVA2:
		return "RVA2: Relative volume adjustment (2)"
	case FrameRVRB:
		return "RVRB: Reverb"
	case FrameSEEK:
		return "SEEK: Seek frame"
	case FrameSIGN:
		return "SIGN: Signature frame"
	case FrameSYLT:
		return "SYLT: Synchronised lyric/text"
	case FrameSYTC:
		return "SYTC: Synchronised tempo codes"
	case FrameTALB:
		return "TALB: Album/Movie/Show title"
	case FrameTBPM:
		return "TBPM: BPM (beats per minute)"
	case FrameTCOM:
		return "TCOM: Composer"
	case FrameTCON:
		return "TCON: Content type"
	case FrameTCOP:
		return "TCOP: Copyright message"
	case FrameTDEN:
		return "TDEN: Encoding time"
	case FrameTDLY:
		return "TDLY: Playlist delay"
	case FrameTDOR:
		return "TDOR: Original release time"
	case FrameTDRC:
		return "TDRC: Recording time"
	case FrameTDRL:
		return "TDRL: Release time"
	case FrameTDTG:
		return "TDTG: Tagging time"
	case FrameTENC:
		return "TENC: Encoded by"
	case FrameTEXT:
		return "TEXT: Lyricist/Text writer"
	case FrameTFLT:
		return "TFLT: File type"
	case FrameTIPL:
		return "TIPL: Involved people list"
	case FrameTIT1:
		return "TIT1: Content group description"
	case FrameTIT2:
		return "TIT2: Title/songname/content description"
	case FrameTIT3:
		return "TIT3: Subtitle/Description refinement"
	case FrameTKEY:
		return "TKEY: Initial key"
	case FrameTLAN:
		return "TLAN: Language(s)"
	case FrameTLEN:
		return "TLEN: Length"
	case FrameTMCL:
		return "TMCL: Musician credits list"
	case FrameTMED:
		return "TMED: Media type"
	case FrameTMOO:
		return "TMOO: Mood"
	case FrameTOAL:
		return "TOAL: Original album/movie/show title"
	case FrameTOFN:
		return "TOFN: Original filename"
	case FrameTOLY:
		return "TOLY: Original lyricist(s)/text writer(s)"
	case FrameTOPE:
		return "TOPE: Original artist(s)/performer(s)"
	case FrameTOWN:
		return "TOWN: File owner/licensee"
	case FrameTPE1:
		return "TPE1: Lead performer(s)/Soloist(s)"
	case FrameTPE2:
		return "TPE2: Band/orchestra/accompaniment"
	case FrameTPE3:
		return "TPE3: Conductor/performer refinement"
	case FrameTPE4:
		return "TPE4: Interpreted, remixed, or otherwise modified by"
	case FrameTPOS:
		return "TPOS: Part of a set"
	case FrameTPRO:
		return "TPRO: Produced notice"
	case FrameTPUB:
		return "TPUB: Publisher"
	case FrameTRCK:
		return "TRCK: Track number/Position in set"
	case FrameTRSN:
		return "TRSN: Internet radio station name"
	case FrameTRSO:
		return "TRSO: Internet radio station owner"
	case FrameTSOA:
		return "TSOA: Album sort order"
	case FrameTSOP:
		return "TSOP: Performer sort order"
	case FrameTSOT:
		return "TSOT: Title sort order"
	case FrameTSRC:
		return "TSRC: ISRC (international standard recording code)"
	case FrameTSSE:
		return "TSSE: Software/Hardware and settings used for encoding"
	case FrameTSST:
		return "TSST: Set subtitle"
	case FrameTXXX:
		return "TXXX: User defined text information frame"
	case FrameUFID:
		return "UFID: Unique file identifier"
	case FrameUSER:
		return "USER: Terms of use"
	case FrameUSLT:
		return "USLT: Unsynchronised lyric/text transcription"
	case FrameWCOM:
		return "WCOM: Commercial information"
	case FrameWCOP:
		return "WCOP: Copyright/Legal information"
	case FrameWOAF:
		return "WOAF: Official audio file webpage"
	case FrameWOAR:
		return "WOAR: Official artist/performer webpage"
	case FrameWOAS:
		return "WOAS: Official audio source webpage"
	case FrameWORS:
		return "WORS: Official Internet radio station homepage"
	case FrameWPAY:
		return "WPAY: Payment"
	case FrameWPUB:
		return "WPUB: Publishers official webpage"
	case FrameWXXX:
		return "WXXX: User defined URL link frame"
	case FrameEQUA:
		return "EQUA: Equalization"
	case FrameIPLS:
		return "IPLS: Involved people list"
	case FrameRVAD:
		return "RVAD: Relative volume adjustment"
	case FrameTDAT:
		return "TDAT: Date"
	case FrameTIME:
		return "TIME: Time"
	case FrameTORY:
		return "TORY: Original release year"
	case FrameTRDA:
		return "TRDA: Recording dates"
	case FrameTSIZ:
		return "TSIZ: Size"
	case FrameTYER:
		return "TYER: Year"
	default:
		buf := [4]byte{
			byte(id >> 24),
			byte(id >> 16),
			byte(id >> 8),
			byte(id),
		}
		return "FrameID(\"" + string(buf[:]) + "\")"
	}
}
