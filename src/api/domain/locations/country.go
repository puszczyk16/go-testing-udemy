package locations

type Country struct {
	Id 				string			`json:"id"`
	Name 			string			`json:"name"`
	TimeZone 		string			`json:"time_zone"`
	GeoInformation 	GeoInformation	`json:"geo_information"`
	States			[]State			`json:"states"`
}

type GeoInformation struct {
	Location GeoLocation	`json:"location"`
}

type GeoLocation struct {
	Latitude	float64 `json:"latitude"`
	Longitude	float64 `json:"longitude"`
}

type State struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
}

/*
https://api.mercadolibre.com/countries/AR

{
	"id": "AR",
	"name": "Argentina",
	"locale": "es_AR",
	"currency_id": "ARS",
	"decimal_separator": ",",
	"thousands_separator": ".",
	"time_zone": "GMT-03:00",
	"geo_information": {
	  "location": {
		"latitude": -38.416096,
		"longitude": -63.616673
	  }
	},
	"states": [
	  {
		"id": "AR-B",
		"name": "Buenos Aires"
	  },
	  {
		"id": "AR-C",
		"name": "Capital Federal"
	  },
	  {
		"id": "AR-K",
		"name": "Catamarca"
	  },
	  {
		"id": "AR-H",
		"name": "Chaco"
	  },
	  {
		"id": "AR-U",
		"name": "Chubut"
	  },
	  {
		"id": "AR-W",
		"name": "Corrientes"
	  },
	  {
		"id": "AR-X",
		"name": "Córdoba"
	  },
	  {
		"id": "AR-E",
		"name": "Entre Ríos"
	  },
	  {
		"id": "AR-P",
		"name": "Formosa"
	  },
	  {
		"id": "AR-Y",
		"name": "Jujuy"
	  },
	  {
		"id": "AR-L",
		"name": "La Pampa"
	  },
	  {
		"id": "AR-F",
		"name": "La Rioja"
	  },
	  {
		"id": "AR-M",
		"name": "Mendoza"
	  },
	  {
		"id": "AR-N",
		"name": "Misiones"
	  },
	  {
		"id": "AR-Q",
		"name": "Neuquén"
	  },
	  {
		"id": "AR-R",
		"name": "Río Negro"
	  },
	  {
		"id": "AR-A",
		"name": "Salta"
	  },
	  {
		"id": "AR-J",
		"name": "San Juan"
	  },
	  {
		"id": "AR-D",
		"name": "San Luis"
	  },
	  {
		"id": "AR-Z",
		"name": "Santa Cruz"
	  },
	  {
		"id": "AR-S",
		"name": "Santa Fe"
	  },
	  {
		"id": "AR-G",
		"name": "Santiago del Estero"
	  },
	  {
		"id": "AR-V",
		"name": "Tierra del Fuego"
	  },
	  {
		"id": "AR-T",
		"name": "Tucumán"
	  }
	]
  }
  */