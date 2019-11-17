/*
This package contains the PhotoPrism REST api.

Additional information can be found in our Developer Guide:

https://github.com/photoprism/photoprism/wiki
*/
package api

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func init() {
	log = logrus.StandardLogger()
}
