// Code generated by generate-tl-files; DO NOT EDIT.

package telegram

type BaseTheme uint32

const (
	BaseThemeArctic  BaseTheme = 0x5b11125a
	BaseThemeClassic BaseTheme = 0xc3a12462
	BaseThemeDay     BaseTheme = 0xfbd81688
	BaseThemeNight   BaseTheme = 0xb7b31ea8
	BaseThemeTinted  BaseTheme = 0x6d5f77ee
)

func (e BaseTheme) String() string {
	switch e {
	case BaseTheme(0x5b11125a):
		return "baseThemeArctic"
	case BaseTheme(0xc3a12462):
		return "baseThemeClassic"
	case BaseTheme(0xfbd81688):
		return "baseThemeDay"
	case BaseTheme(0xb7b31ea8):
		return "baseThemeNight"
	case BaseTheme(0x6d5f77ee):
		return "baseThemeTinted"
	default:
		return "<UNKNOWN BaseTheme>"
	}
}

func (e BaseTheme) CRC() uint32 { return uint32(e) }

type InlineQueryPeerType uint32

const (
	InlineQueryPeerTypeBroadcast InlineQueryPeerType = 0x6334ee9a
	InlineQueryPeerTypeChat      InlineQueryPeerType = 0xd766c50a
	InlineQueryPeerTypeMegagroup InlineQueryPeerType = 0x5ec4be43
	InlineQueryPeerTypePm        InlineQueryPeerType = 0x833c0fac
	InlineQueryPeerTypeSameBotPm InlineQueryPeerType = 0x3081ed9d
)

func (e InlineQueryPeerType) String() string {
	switch e {
	case InlineQueryPeerType(0x6334ee9a):
		return "inlineQueryPeerTypeBroadcast"
	case InlineQueryPeerType(0xd766c50a):
		return "inlineQueryPeerTypeChat"
	case InlineQueryPeerType(0x5ec4be43):
		return "inlineQueryPeerTypeMegagroup"
	case InlineQueryPeerType(0x833c0fac):
		return "inlineQueryPeerTypePM"
	case InlineQueryPeerType(0x3081ed9d):
		return "inlineQueryPeerTypeSameBotPM"
	default:
		return "<UNKNOWN InlineQueryPeerType>"
	}
}

func (e InlineQueryPeerType) CRC() uint32 { return uint32(e) }

type InputPrivacyKey uint32

const (
	InputPrivacyKeyAddedByPhone    InputPrivacyKey = 0xd1219bdd
	InputPrivacyKeyChatInvite      InputPrivacyKey = 0xbdfb0426
	InputPrivacyKeyForwards        InputPrivacyKey = 0xa4dd4c08
	InputPrivacyKeyPhoneCall       InputPrivacyKey = 0xfabadc5f
	InputPrivacyKeyPhoneNumber     InputPrivacyKey = 0x352dafa
	InputPrivacyKeyPhoneP2P        InputPrivacyKey = 0xdb9e70d2
	InputPrivacyKeyProfilePhoto    InputPrivacyKey = 0x5719bacc
	InputPrivacyKeyStatusTimestamp InputPrivacyKey = 0x4f96cb18
)

func (e InputPrivacyKey) String() string {
	switch e {
	case InputPrivacyKey(0xd1219bdd):
		return "inputPrivacyKeyAddedByPhone"
	case InputPrivacyKey(0xbdfb0426):
		return "inputPrivacyKeyChatInvite"
	case InputPrivacyKey(0xa4dd4c08):
		return "inputPrivacyKeyForwards"
	case InputPrivacyKey(0xfabadc5f):
		return "inputPrivacyKeyPhoneCall"
	case InputPrivacyKey(0x352dafa):
		return "inputPrivacyKeyPhoneNumber"
	case InputPrivacyKey(0xdb9e70d2):
		return "inputPrivacyKeyPhoneP2P"
	case InputPrivacyKey(0x5719bacc):
		return "inputPrivacyKeyProfilePhoto"
	case InputPrivacyKey(0x4f96cb18):
		return "inputPrivacyKeyStatusTimestamp"
	default:
		return "<UNKNOWN InputPrivacyKey>"
	}
}

func (e InputPrivacyKey) CRC() uint32 { return uint32(e) }

type PhoneCallDiscardReason uint32

const (
	PhoneCallDiscardReasonBusy       PhoneCallDiscardReason = 0xfaf7e8c9
	PhoneCallDiscardReasonDisconnect PhoneCallDiscardReason = 0xe095c1a0
	PhoneCallDiscardReasonHangup     PhoneCallDiscardReason = 0x57adc690
	PhoneCallDiscardReasonMissed     PhoneCallDiscardReason = 0x85e42301
)

func (e PhoneCallDiscardReason) String() string {
	switch e {
	case PhoneCallDiscardReason(0xfaf7e8c9):
		return "phoneCallDiscardReasonBusy"
	case PhoneCallDiscardReason(0xe095c1a0):
		return "phoneCallDiscardReasonDisconnect"
	case PhoneCallDiscardReason(0x57adc690):
		return "phoneCallDiscardReasonHangup"
	case PhoneCallDiscardReason(0x85e42301):
		return "phoneCallDiscardReasonMissed"
	default:
		return "<UNKNOWN PhoneCallDiscardReason>"
	}
}

func (e PhoneCallDiscardReason) CRC() uint32 { return uint32(e) }

type PrivacyKey uint32

const (
	PrivacyKeyAddedByPhone    PrivacyKey = 0x42ffd42b
	PrivacyKeyChatInvite      PrivacyKey = 0x500e6dfa
	PrivacyKeyForwards        PrivacyKey = 0x69ec56a3
	PrivacyKeyPhoneCall       PrivacyKey = 0x3d662b7b
	PrivacyKeyPhoneNumber     PrivacyKey = 0xd19ae46d
	PrivacyKeyPhoneP2P        PrivacyKey = 0x39491cc8
	PrivacyKeyProfilePhoto    PrivacyKey = 0x96151fed
	PrivacyKeyStatusTimestamp PrivacyKey = 0xbc2eab30
)

func (e PrivacyKey) String() string {
	switch e {
	case PrivacyKey(0x42ffd42b):
		return "privacyKeyAddedByPhone"
	case PrivacyKey(0x500e6dfa):
		return "privacyKeyChatInvite"
	case PrivacyKey(0x69ec56a3):
		return "privacyKeyForwards"
	case PrivacyKey(0x3d662b7b):
		return "privacyKeyPhoneCall"
	case PrivacyKey(0xd19ae46d):
		return "privacyKeyPhoneNumber"
	case PrivacyKey(0x39491cc8):
		return "privacyKeyPhoneP2P"
	case PrivacyKey(0x96151fed):
		return "privacyKeyProfilePhoto"
	case PrivacyKey(0xbc2eab30):
		return "privacyKeyStatusTimestamp"
	default:
		return "<UNKNOWN PrivacyKey>"
	}
}

func (e PrivacyKey) CRC() uint32 { return uint32(e) }

type ReportReason uint32

const (
	InputReportReasonChildAbuse    ReportReason = 0xadf44ee3
	InputReportReasonCopyright     ReportReason = 0x9b89f93a
	InputReportReasonFake          ReportReason = 0xf5ddd6e7
	InputReportReasonGeoIrrelevant ReportReason = 0xdbd4feed
	InputReportReasonOther         ReportReason = 0xc1e4a2b1
	InputReportReasonPornography   ReportReason = 0x2e59d922
	InputReportReasonSpam          ReportReason = 0x58dbcab8
	InputReportReasonViolence      ReportReason = 0x1e22c78d
)

func (e ReportReason) String() string {
	switch e {
	case ReportReason(0xadf44ee3):
		return "inputReportReasonChildAbuse"
	case ReportReason(0x9b89f93a):
		return "inputReportReasonCopyright"
	case ReportReason(0xf5ddd6e7):
		return "inputReportReasonFake"
	case ReportReason(0xdbd4feed):
		return "inputReportReasonGeoIrrelevant"
	case ReportReason(0xc1e4a2b1):
		return "inputReportReasonOther"
	case ReportReason(0x2e59d922):
		return "inputReportReasonPornography"
	case ReportReason(0x58dbcab8):
		return "inputReportReasonSpam"
	case ReportReason(0x1e22c78d):
		return "inputReportReasonViolence"
	default:
		return "<UNKNOWN ReportReason>"
	}
}

func (e ReportReason) CRC() uint32 { return uint32(e) }

type SecureValueType uint32

const (
	SecureValueTypeAddress               SecureValueType = 0xcbe31e26
	SecureValueTypeBankStatement         SecureValueType = 0x89137c0d
	SecureValueTypeDriverLicense         SecureValueType = 0x6e425c4
	SecureValueTypeEmail                 SecureValueType = 0x8e3ca7ee
	SecureValueTypeIdentityCard          SecureValueType = 0xa0d0744b
	SecureValueTypeInternalPassport      SecureValueType = 0x99a48f23
	SecureValueTypePassport              SecureValueType = 0x3dac6a00
	SecureValueTypePassportRegistration  SecureValueType = 0x99e3806a
	SecureValueTypePersonalDetails       SecureValueType = 0x9d2a81e3
	SecureValueTypePhone                 SecureValueType = 0xb320aadb
	SecureValueTypeRentalAgreement       SecureValueType = 0x8b883488
	SecureValueTypeTemporaryRegistration SecureValueType = 0xea02ec33
	SecureValueTypeUtilityBill           SecureValueType = 0xfc36954e
)

func (e SecureValueType) String() string {
	switch e {
	case SecureValueType(0xcbe31e26):
		return "secureValueTypeAddress"
	case SecureValueType(0x89137c0d):
		return "secureValueTypeBankStatement"
	case SecureValueType(0x6e425c4):
		return "secureValueTypeDriverLicense"
	case SecureValueType(0x8e3ca7ee):
		return "secureValueTypeEmail"
	case SecureValueType(0xa0d0744b):
		return "secureValueTypeIdentityCard"
	case SecureValueType(0x99a48f23):
		return "secureValueTypeInternalPassport"
	case SecureValueType(0x3dac6a00):
		return "secureValueTypePassport"
	case SecureValueType(0x99e3806a):
		return "secureValueTypePassportRegistration"
	case SecureValueType(0x9d2a81e3):
		return "secureValueTypePersonalDetails"
	case SecureValueType(0xb320aadb):
		return "secureValueTypePhone"
	case SecureValueType(0x8b883488):
		return "secureValueTypeRentalAgreement"
	case SecureValueType(0xea02ec33):
		return "secureValueTypeTemporaryRegistration"
	case SecureValueType(0xfc36954e):
		return "secureValueTypeUtilityBill"
	default:
		return "<UNKNOWN SecureValueType>"
	}
}

func (e SecureValueType) CRC() uint32 { return uint32(e) }

type TopPeerCategory uint32

const (
	TopPeerCategoryBotsInline     TopPeerCategory = 0x148677e2
	TopPeerCategoryBotsPm         TopPeerCategory = 0xab661b5b
	TopPeerCategoryChannels       TopPeerCategory = 0x161d9628
	TopPeerCategoryCorrespondents TopPeerCategory = 0x637b7ed
	TopPeerCategoryForwardChats   TopPeerCategory = 0xfbeec0f0
	TopPeerCategoryForwardUsers   TopPeerCategory = 0xa8406ca9
	TopPeerCategoryGroups         TopPeerCategory = 0xbd17a14a
	TopPeerCategoryPhoneCalls     TopPeerCategory = 0x1e76a78c
)

func (e TopPeerCategory) String() string {
	switch e {
	case TopPeerCategory(0x148677e2):
		return "topPeerCategoryBotsInline"
	case TopPeerCategory(0xab661b5b):
		return "topPeerCategoryBotsPM"
	case TopPeerCategory(0x161d9628):
		return "topPeerCategoryChannels"
	case TopPeerCategory(0x637b7ed):
		return "topPeerCategoryCorrespondents"
	case TopPeerCategory(0xfbeec0f0):
		return "topPeerCategoryForwardChats"
	case TopPeerCategory(0xa8406ca9):
		return "topPeerCategoryForwardUsers"
	case TopPeerCategory(0xbd17a14a):
		return "topPeerCategoryGroups"
	case TopPeerCategory(0x1e76a78c):
		return "topPeerCategoryPhoneCalls"
	default:
		return "<UNKNOWN TopPeerCategory>"
	}
}

func (e TopPeerCategory) CRC() uint32 { return uint32(e) }

type AuthCodeType uint32

const (
	AuthCodeTypeCall      AuthCodeType = 0x741cd3e3
	AuthCodeTypeFlashCall AuthCodeType = 0x226ccefb
	AuthCodeTypeSms       AuthCodeType = 0x72a3158c
)

func (e AuthCodeType) String() string {
	switch e {
	case AuthCodeType(0x741cd3e3):
		return "auth.codeTypeCall"
	case AuthCodeType(0x226ccefb):
		return "auth.codeTypeFlashCall"
	case AuthCodeType(0x72a3158c):
		return "auth.codeTypeSms"
	default:
		return "<UNKNOWN auth.CodeType>"
	}
}

func (e AuthCodeType) CRC() uint32 { return uint32(e) }

type StorageFileType uint32

const (
	StorageFileGif     StorageFileType = 0xcae1aadf
	StorageFileJpeg    StorageFileType = 0x7efe0e
	StorageFileMov     StorageFileType = 0x4b09ebbc
	StorageFileMp3     StorageFileType = 0x528a0677
	StorageFileMp4     StorageFileType = 0xb3cea0e4
	StorageFilePartial StorageFileType = 0x40bc6f52
	StorageFilePdf     StorageFileType = 0xae1e508d
	StorageFilePng     StorageFileType = 0xa4f63c0
	StorageFileUnknown StorageFileType = 0xaa963b05
	StorageFileWebp    StorageFileType = 0x1081464c
)

func (e StorageFileType) String() string {
	switch e {
	case StorageFileType(0xcae1aadf):
		return "storage.fileGif"
	case StorageFileType(0x7efe0e):
		return "storage.fileJpeg"
	case StorageFileType(0x4b09ebbc):
		return "storage.fileMov"
	case StorageFileType(0x528a0677):
		return "storage.fileMp3"
	case StorageFileType(0xb3cea0e4):
		return "storage.fileMp4"
	case StorageFileType(0x40bc6f52):
		return "storage.filePartial"
	case StorageFileType(0xae1e508d):
		return "storage.filePdf"
	case StorageFileType(0xa4f63c0):
		return "storage.filePng"
	case StorageFileType(0xaa963b05):
		return "storage.fileUnknown"
	case StorageFileType(0x1081464c):
		return "storage.fileWebp"
	default:
		return "<UNKNOWN storage.FileType>"
	}
}

func (e StorageFileType) CRC() uint32 { return uint32(e) }
