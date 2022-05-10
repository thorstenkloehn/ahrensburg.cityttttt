function onEachFeature(feature, layer) {
    // does this feature have a property named popupContent?
    if (feature.properties && feature.properties.popupContent) {
        layer.bindPopup(feature.properties.popupContent);
    }
}
//<![CDATA[

var map = L.map('map').setView([53.6700755, 10.2071975], 13);

L.tileLayer('https://ahrensburg.city/karte/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://ahrensburg.city">ahrensburg.city</a>,<a href="./docs/Impressum.html">Impressum</a>,<a href="./docs/Datenschutzerklarung">Datenschutzerklärung</a> , </a><a href="./docs/">Dokument</a>,&copy;<a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);



let laden = new XMLHttpRequest;
laden.open('GET', 'geoserver/ahrensburg_city/ows?service=WFS&version=1.0.0&request=GetFeature&typeName=ahrensburg_city%3AKartendaten&maxFeatures=50&outputFormat=application%2Fjson', true);
laden.onload = function () {
    if (laden.status == 200) {

        L.geoJSON(JSON.parse(laden.responseText), {
            onEachFeature: onEachFeature
        }).addTo(map);

    }
}

laden.send();

// test

