CREATE TABLE IF NOT EXISTS linear_mortgage (
	"type" TEXT NOT NULL,
	"mortgage" INTEGER NOT NULL,
    "terms" INT NOT NULL,
	"interest" FLOAT NOT NULL,
	"month" INTEGER,
	"year" INTEGER,
	"remainder" FLOAT,
	"interest_pmt" FLOAT,
	"redemption_pmt" FLOAT,
	"total_pmt" FLOAT,
    "tax_discount" FLOAT,
    "nett_pmt" FLOAT,
    "payed_interest" FLOAT,
    "payed_redemption" FLOAT,
    "total_payed" FLOAT,
    "nett_payed" FLOAT
);
