package pdfreader

import (
    "github.com/pdfcpu/pdfcpu/pkg/api" // Make sure this is correctly imported
    "os"
    "strings"
)

// ReadPDFText extracts text from a given PDF file path
func ReadPDFText(filePath string) (string, error) {
    // Open the PDF file
    f, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer f.Close()

    var sb strings.Builder  // StringBuilder to collect text
    ctx, err := api.ReadContextFile(filePath)
    if err != nil {
        return "", err
    }

    for _, page := range ctx.Pages {
        if page.V.IsNull() {
            continue
        }
        extractText, err := api.ExtractPageText(page)
        if err != nil {
            return "", err
        }
        sb.WriteString(extractText.Text)
    }

    return sb.String(), nil
}
