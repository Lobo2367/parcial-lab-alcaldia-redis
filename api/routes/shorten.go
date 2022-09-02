package routes

import()

type request struct{
	URL
	CustomShort
	Expiry
}

type response struct{
	URL
	CustomShort
	Expiry
	XRateRemaining
	XRateLimitRest
}
