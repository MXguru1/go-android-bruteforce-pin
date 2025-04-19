package hid

var TouchscreenReportDesc = []byte{
	0x05, 0x0D, // Usage Page (Digitizer)
	0x09, 0x02, // Usage (Pen)
	0xA1, 0x01, // Collection (Application)
	0x09, 0x20, //   Usage (Stylus)
	0xA1, 0x00, //   Collection (Physical)
	0x09, 0x42, //     Usage (Tip Switch)
	0x09, 0x32, //     Usage (In Range)
	0x15, 0x00, //     Logical Minimum (0)
	0x25, 0x01, //     Logical Maximum (1)
	0x75, 0x01, //     Report Size (1)
	0x95, 0x02, //     Report Count (2)
	0x81, 0x02, //     Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x75, 0x01, //     Report Size (1)
	0x95, 0x06, //     Report Count (6)
	0x81, 0x01, //     Input (Const,Array,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x05, 0x01, //     Usage Page (Generic Desktop Ctrls)
	0x09, 0x01, //     Usage (Pointer)
	0xA1, 0x00, //     Collection (Physical)
	0x09, 0x30, //       Usage (X)
	0x09, 0x31, //       Usage (Y)
	0x16, 0x00, 0x00, //       Logical Minimum (0)
	0x26, 0x10, 0x27, //       Logical Maximum (10000)
	0x36, 0x00, 0x00, //       Physical Minimum (0)
	0x46, 0x10, 0x27, //       Physical Maximum (10000)
	0x66, 0x00, 0x00, //       Unit (None)
	0x75, 0x10, //       Report Size (16)
	0x95, 0x02, //       Report Count (2)
	0x81, 0x02, //       Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x91, 0x02, //       Output (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position,Non-volatile)
	0x95, 0x03, //       Report Count (3)
	0x91, 0x02, //       Output (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position,Non-volatile)
	0xC0, //     End Collection
	0xC0, //   End Collection
	0xC0, // End Collection

	// 70 bytes

}

var TouchscreenReportDesc2 = []byte{
	0x05, 0x0d, // USAGE_PAGE (Digitizers)
	0x09, 0x04, // USAGE (Touch Screen)
	0xa1, 0x01, // COLLECTION (Application)
	0x09, 0x22, //   USAGE (Finger)
	0xa1, 0x00, //   COLLECTION (Physical)
	0x09, 0x42, //     USAGE (Tip Switch)
	0x09, 0x51, //     USAGE (Contact Identifier) *added line
	0x15, 0x00, //     LOGICAL_MINIMUM (0)
	0x25, 0x01, //     LOGICAL_MAXIMUM (1)
	0x75, 0x01, //     REPORT_SIZE (1)
	0x95, 0x02, //     REPORT_COUNT (2)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0x95, 0x0e, //     REPORT_COUNT (14)
	0x81, 0x03, //     INPUT (Cnst,Var,Abs)
	0x05, 0x01, //     USAGE_PAGE (Generic Desktop)
	0x75, 0x10, //     REPORT_SIZE (16)
	0x95, 0x01, //     REPORT_COUNT (1)
	0x55, 0x0d, //     UNIT_EXPONENT (-3)
	0x65, 0x33, //     UNIT (Inch,EngLinear)
	0x15, 0x00, //     LOGICAL_MINIMUM (0)
	0x26, 0xff, 0x7f, //     LOGICAL_MAXIMUM (32767)
	0x09, 0x30, //     USAGE (X)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0x09, 0x31, //     USAGE (Y)
	0x81, 0x02, //     INPUT (Data,Var,Abs)
	0xc0, //   END_COLLECTION
	0xc0, // END_COLLECTION
}
