package recover

import (
	"regexp"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func New() fiber.Handler {
	return recover.New(
		recover.Config{
			EnableStackTrace: true,
			StackTraceHandler: func(c fiber.Ctx, e any) {
				buf := make([]byte, 1024)
				buf = buf[:runtime.Stack(buf, false)]

				var messages []string

				if compile, err := regexp.Compile(`(^.*:\d+\s+)`); err == nil {
					for _, x := range strings.Split(string(buf), "\n") {
						if v := compile.FindString(x); v != "" {
							if strings.Contains(v, "/pkg/mod/") {
								continue
							}

							if strings.Contains(v, "/go/src/") {
								continue
							}

							messages = append(messages, strings.TrimSpace(v))
						}
					}
				}

				if len(messages) > 1 {
					messages = messages[1:]
				}

				c.Locals("stack", messages)
			},
		},
	)
}
