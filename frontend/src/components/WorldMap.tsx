import Map from "react-map-gl/maplibre";
import "maplibre-gl/dist/maplibre-gl.css";

const MAP_STYLE = "https://tiles.openfreemap.org/styles/liberty";

function WorldMap() {
  return (
    <Map
      initialViewState={{
        longitude: -100,
        latitude: 60,
        zoom: 3.5,
      }}
      minZoom={1}
      dragRotate={false}
      mapStyle={MAP_STYLE}
      style={{ width: "100%", height: "100%" }}
    />
  );
}

export default WorldMap;
