<h1 align="center" style="margin-top: 0px;">iap</h1>

<p align="center">
<a href="https://github.com/space-code/iap/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/space-code/iap?style=flat"></a> 
<a href="https://github.com/space-code/iap"><img alt="CI" src="https://github.com/space-code/iap/actions/workflows/ci.yml/badge.svg?branch=main"></a>
</p>

## Description
`iap` verifies a purchase receipt via AppStore

- [Usage](#usage)
- [Communication](#communication)
- [Contributing](#contributing)
- [Author](#author)
- [License](#license)

## Usage

```go
package appstore

// create an IAP client
client := appstore.New()
// create an validation request
r := appstore.IAPValidationRequest{
  ReceiptData:            "enter your receipt here",
  ExcludeOldTransactions: true,
}
ctx := context.Background()
var response appstore.IAPValidationResponse
// receipt validation
err := v.VerifyReceipt(ctx, req, &response)
```

## Communication
- If you **found a bug**, open an issue.
- If you **have a feature request**, open an issue.
- If you **want to contribute**, submit a pull request.

## Contributing
Please feel free to help out with this project! If you see something that could be made better or want a new feature, open up an issue or send a Pull Request!

## Author
Nikita Vasilev, nv3212@gmail.com

## License
iap is available under the MIT license. See the LICENSE file for more info.