mapboxgl.accessToken = 'tokentokentoken';

const coords = [13.411975805000539, 52.53134099392974];

const map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/outdoors-v10?optimize=true',
    center: coords,
    zoom: 13
});

map.on('load', () => {
    // Add a source and layer displaying a point which will be animated in a circle.
    map.addSource('point', {
        type: 'geojson',
        data: {
            type: 'Point',
            coordinates: coords,
        }
    });

    const points = [
        { name: 'outline', radius: 13, color: '#ffffff' },
        { name: 'color',   radius: 10, color: '#00ffcc' },
        { name: 'dot',     radius: 4,  color: '#444444' },
    ];

    points.forEach(point => {
        map.addLayer({
            id: point.name,
            source: 'point',
            type: 'circle',
            paint: {
                'circle-radius': point.radius,
                'circle-color': point.color
            }
        });
    });
});
