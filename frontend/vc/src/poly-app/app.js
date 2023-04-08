class StaticGlobals {
  //const HOST = "sasha123.ddns.ukrtel.net:1234";

  constructor() {
    console.log("constants initiates");
  };

  static GetHostName() {
    var host = window.location.hostname
    var port = window.location.port
    if (host == "localhost") {
      host = "sasha123.ddns.ukrtel.net:1234";
    } else {
      host = host + ":" + port;
    }
    console.log("host:" + host);
    return host;
  };

};
