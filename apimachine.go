package apimachine

import (
        "bytes"
        "encoding/json"
        "errors"
        "net/http"
        "time"
)

const baseURL = "https://api.github.com"
