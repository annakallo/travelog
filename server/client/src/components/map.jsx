import { VectorMap } from "react-jvectormap"
import React from "react";

const mapData = {
    CN: 100000,
    IT: 0,
    GB: 0,
    NL: 0,
    FR: 0,
    ES: 0,
    RO: 0,
    HU: 0,
};


const handleClick = (e, countryCode) => {
    console.log(countryCode);
};

const Map = () => {
    return (
        <div>
            <VectorMap
                map={"world_mill"}
                backgroundColor="transparent" //change it to ocean blue: #0077be
                zoomOnScroll={true}
                containerStyle={{
                    width: "100%",
                    height: "520px"
                }}
                onRegionClick={handleClick} //gets the country code
                containerClassName="map"
                regionStyle={{
                    initial: {
                        fill: "#e4e4e4",
                        "fill-opacity": 0.9,
                        stroke: "none",
                        "stroke-width": 0,
                        "stroke-opacity": 0
                    },
                    hover: {
                        "fill-opacity": 0.8,
                        cursor: "pointer"
                    },
                    selected: {
                        fill: "#146804" //color for the clicked country
                    },
                    selectedHover: {}
                }}
                regionsSelectable={true}
                series={{
                    regions: [
                        {
                            values: mapData, //this is your data
                            scale: ["#2938bc", "#e4e4e4" ], //your color game's here
                            normalizeFunction: "polynomial"
                        }
                    ]
                }}
            />
        </div>
    );
};
export default Map;
