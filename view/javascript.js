function onEachFeature(feature, layer) {
    // does this feature have a property named popupContent?
    if (feature.properties && feature.properties.popupContent) {
        layer.bindPopup(feature.properties.popupContent);
    }
}
//<![CDATA[

var map = L.map('map').setView([53.6700755, 10.2071975], 13);

L.tileLayer('https://ahrensburg.city/karte/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://ahrensburg.city">ahrensburg.city</a>,<a href="./docs/Impressum">Impressum</a>,<a href="./docs/Datenschutzerklärung">Datenschutzerklärung</a> , </a><a href="./docs/">Dokument</a>,&copy;<a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);



let laden = new XMLHttpRequest;
laden.open('GET', 'geoserver/Ahrensburg/ows?service=WFS&version=1.0.0&request=GetFeature&typeName=Ahrensburg%3Aahrensburg&maxFeatures=50&outputFormat=application%2Fjson', true);
laden.onload = function () {
    if (laden.status == 200) {

        L.geoJSON(JSON.parse(laden.responseText), {
            onEachFeature: onEachFeature
        }).addTo(map);

    }
}

laden.send();

var gpx = './gpx/Auewanderweg/auewanderweg_gpx_20220519_191032.gpx'; // URL to your GPX file or the GPX itself

new L.GPX(gpx, {
    async: true,
    marker_options: {
        startIconUrl: 'static/Leaflet/gpx/pin-icon-start.png',
        endIconUrl: 'static/Leaflet/gpx/pin-icon-end.png',
        shadowUrl: 'static/Leaflet/gpx/pin-shadow.png',
        wptIconUrls: 'static/Leaflet/gpx/pin-icon-wpt.png'
    }
}).on('loaded', function(e) {
    map.fitBounds(e.target.getBounds());
}).addTo(map);