package regmap

// DeviceESM4450 is taken from the EMKO protocol documentation.  The descriptions and
// some of the comments are from the original documentation and are left more or less
// in their original form -- with its charming attempt at english.
var DeviceESM4450 = Device{
	Name:        "ESM-4450",
	Description: "PID controller from EMKO",
	Registers: []Register{
		// Operator parameters
		// Set list (SEt LySt)
		{
			Address:     40001,
			Symbol:      "process_sv",
			Display:     "PSEt",
			Description: "SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40002,
			Symbol:      "alarm_1_sv",
			Display:     "ALr1",
			Description: "if process input selected SU-L SU-i, if analog input selected SUL2-SUu2",
			Access:      ModeReadWrite,
		},
		{
			Address:     40003,
			Symbol:      "alarm_3_sv",
			Display:     "ALr3",
			Description: "if process input selected SU-L SU-i, if analog input selected SUL2-SUu2",
			Access:      ModeReadWrite,
		},
		// Running mode (run LySt)
		{
			Address:     40005,
			Symbol:      "tuning_type",
			Display:     "tunn",
			Description: "no, Atun, Stun, At.St",
			Access:      ModeReadWrite,
		},
		{
			Address:     40006,
			Symbol:      "auto_tuning",
			Display:     "Attn",
			Description: "no YES",
			Access:      ModeReadWrite,
		},
		{
			Address:     40007,
			Symbol:      "control_input_auto",
			Display:     "Auto",
			Description: "Manual/Automatic selection for control input",
			Access:      ModeReadWrite,
		},
		{
			Address:     40008,
			Symbol:      "ramp_soak_selection",
			Display:     "rSSL",
			Description: "Off, run, HoLd",
			Access:      ModeReadWrite,
		},
		{
			Address:     40009,
			Symbol:      "bumpless_transfer",
			Display:     "bPLt",
			Description: "no, YES",
			Access:      ModeReadWrite,
		},
		{
			Address: 40167,
			Symbol:  "valve_control_type",
			Display: "ULSL",
			Description: `
			valve control type selection

			(CAUTION!!When you are changing this parameter there must be no electrical connections.)

			(if module 1 is relay out you can see and change this parameter else you can't use valve control)
			0 = no valve control
			1 = heating (reverse action)
			2 = cooling (direct action)
			if your choice is heating or cooling;
			-modul1 output uses for open the valve and out3 output for close the
			valve.
			-you can't select pid output at the modul2 output)
			`,
			Access: ModeReadWrite,
		},

		// Display list (dySP LySt)
		{
			Address: 40010,
			Symbol:  "top_display",
			Display: "tdSP",
			Description: `
			0 PV
 			1 Deviation (SV-PV)
 			2 2nd sensor input (if equipment has 2nd sensor module)
			`,
			Access: ModeReadWrite,
		},
		{
			Address: 40011,
			Symbol:  "bottom_display",
			Display: "bdSP",
			Description: `
			Bottom display (if working man mode always shows Power)
			
			0 Local SV
			1 Power(%)
			2 ramp - soak display
			- no ramp - soak P.End
			- ramp - soak waiting HoLd
			- ramp segment rA 1-8
			- soak segment So 1-8
			3 2nd sensor input (if equipment has 2nd sensor module)
			`,
			Access: ModeReadWrite,
		},

		// Ramp/Soak (rmP SoA)
		{
			Address: 40012,
			Symbol:  "start_ramp_time",
			Display: "STrA",
			Description: `
			0 to 99h 59min
 			if set 0.0 start ramp doesn’t work
			`,
			Access: ModeReadWrite,
		},
		{
			Address: 40013,
			Symbol:  "ramp_soak_tolerance",
			Display: "rSto",
			Description: `
			0 - %50 F.S.
			if 0 doesn’t work
 			if (set value+rSto)<temp <(set value+rSto) ramp or soak time
 			working else ramp or soak time holding & process waiting to came
 			normal position.
			`,
			Access: ModeReadWrite,
		},
		{
			Address: 40014,
			Symbol:  "ramp_soak_type",
			Display: "rStY",
			Description: `
			0 1-4 segment
 			1 5-8 segment
 			2 1-8 segment
			`,
			Access: ModeReadWrite,
		},
		{
			Address:     40015,
			Symbol:      "target_sv_1",
			Display:     "PU_1",
			Description: "SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40016,
			Symbol:      "ramp_segment_time_1",
			Display:     "tr_1",
			Description: "1. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40017,
			Symbol:      "soak_segment_time_1",
			Display:     "tS_1",
			Description: "1. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40018,
			Symbol:      "target_sv_2",
			Display:     "PU_2",
			Description: "2. target SV SU-L – SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40019,
			Symbol:      "ramp_segment_time_2",
			Display:     "tr_2",
			Description: "2. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40020,
			Symbol:      "soak_segment_time_2",
			Display:     "tS_2",
			Description: "2. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40021,
			Symbol:      "target_sv_3",
			Display:     "PU_3",
			Description: "3. target SV SU-L – SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40022,
			Symbol:      "ramp_segment_time_3",
			Display:     "tr_3",
			Description: "3. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40023,
			Symbol:      "soak_segment_time_3",
			Display:     "tS_3",
			Description: "3. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40024,
			Symbol:      "target_sv_4",
			Display:     "PU_4",
			Description: "4. target SV SU-L – SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40025,
			Symbol:      "ramp_segment_time_4",
			Display:     "tr_4",
			Description: "4. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40026,
			Symbol:      "soak_segment_time_4",
			Display:     "tS_4",
			Description: "4. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40027,
			Symbol:      "target_sv_5",
			Display:     "PU_5",
			Description: "5. target SV SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40028,
			Symbol:      "ramp_segment_time_5",
			Display:     "tr_5",
			Description: "5. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40029,
			Symbol:      "soak_segment_time_5",
			Display:     "tS_5",
			Description: "5. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40030,
			Symbol:      "target_sv_6",
			Display:     "PU_6",
			Description: "6. target SV SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40031,
			Symbol:      "ramp_segment_time_6",
			Display:     "tr_6",
			Description: "6. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40032,
			Symbol:      "soak_segment_time_6",
			Display:     "tS_6",
			Description: "6. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40033,
			Symbol:      "target_sv_7",
			Display:     "PU_7",
			Description: "7. target SV SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40034,
			Symbol:      "ramp_segment_time_7",
			Display:     "tr_7",
			Description: "7. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40035,
			Symbol:      "soak_segment_time_7",
			Display:     "tS_7",
			Description: "7. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40036,
			Symbol:      "target_sv_8",
			Display:     "PU_8",
			Description: "8. target SV SU-L - SU-u",
			Access:      ModeReadWrite,
		},
		{
			Address:     40037,
			Symbol:      "ramp_segment_time_8",
			Display:     "tr_8",
			Description: "8. ramp segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},
		{
			Address:     40038,
			Symbol:      "soak_segment_time_8",
			Display:     "tS_8",
			Description: "8. soak segment time 0 to 99h 59min",
			Access:      ModeReadWrite,
		},

		// Technician parameters
		{
			Address:     40065,
			Symbol:      "input_range_lower_limit",
			Display:     "LoL",
			Description: "scale min - upL",
			Access:      ModeReadWrite,
		},
		{
			Address:     40066,
			Symbol:      "input_range_upper_limit",
			Display:     "uPL",
			Description: "LoL - scale max (FS=upL - LoL)",
			Access:      ModeReadWrite,
		},
		{
			Address:     40067,
			Symbol:      "pv_offset",
			Display:     "PuoF",
			Description: " -10 to 10% FS, With this function, predetermined value is added to the input reading",
			Access:      ModeReadWrite,
		},
		{
			Address:     40068,
			Symbol:      "input_filter_time_constant",
			Display:     "yFLt",
			Description: "0.0 to 900.0 seconds",
			Access:      ModeReadWrite,
		},
		{
			Address:     40069,
			Symbol:      "cold_junction",
			Display:     "CJnC",
			Description: "no = does not perform the RCJ, yes = performs the RCJ",
			Access:      ModeReadWrite,
		},

		// Process input configuration

	},
}
