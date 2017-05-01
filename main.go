package main

import (
	"fmt"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/unixpickle/cuda"
	"github.com/unixpickle/essentials"
)

const Indent = "  "

func main() {
	devs, err := cuda.AllDevices()
	must(err)
	for _, dev := range devs {
		fmt.Println()

		var fields [][2]string
		addField := func(name, val string) {
			fields = append(fields, [2]string{name, val})
		}

		name, err := dev.Name()
		must(err)
		addField("Name", name)

		ctx, err := cuda.NewContext(dev, 1)
		must(err)

		must(<-ctx.Run(func() error {
			free, total, err := cuda.MemInfo()
			if err != nil {
				return err
			}
			addField("Total memory", humanize.Bytes(total))
			addField("Free memory", humanize.Bytes(free))
			return nil
		}))

		clock, err := dev.Attr(cuda.DevAttrClockRate)
		must(err)
		addField("Clock speed", formatKHz(clock))

		clock, err = dev.Attr(cuda.DevAttrMemoryClockRate)
		must(err)
		addField("Memory clock", formatKHz(clock))

		// Attribute not supported on CUDA 7.5.
		ratio, err := dev.Attr(cuda.DevAttrSingleToDoublePrecisionPerfRatio)
		if err == nil {
			addField("Single/double perf", strconv.Itoa(ratio))
		}

		printFields(fields)
	}

	fmt.Println()
}

func printFields(fields [][2]string) {
	maxLen := 0
	for _, f := range fields {
		maxLen = essentials.MaxInt(maxLen, len(f[0]))
	}

	for _, f := range fields {
		name, value := f[0], f[1]
		for len(name) < maxLen {
			name = " " + name
		}
		fmt.Println(Indent + name + ": " + value)
	}
}

func formatKHz(val int) string {
	res := humanize.Bytes(uint64(val) * 1000)
	return strings.Replace(res, "B", "Hz", 1)
}

func must(err error) {
	if err != nil {
		essentials.Die(err)
	}
}
