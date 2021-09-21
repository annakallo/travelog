import Highcharts from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import React from "react";
import highchartsMap from "highcharts/modules/map";
import proj4 from "proj4";
import mapDataWorld from "@highcharts/map-collection/custom/world-highres3.geo.json";

highchartsMap(Highcharts);

if (typeof window !== "undefined") {
    window.proj4 = window.proj4 || proj4;
}

const MapHCH = (props) => {
    const mapOptions = {
        chart: {
            map: 'custom/world-highres3',
            borderWidth: 1
        },
        title: {
            text: '2022 Edition'
        },
        colors: ['rgba(19,64,117,0.05)', 'rgba(19,64,117,0.2)', 'rgba(19,64,117,0.4)',
            'rgba(19,64,117,0.5)', 'rgba(19,64,117,0.6)', 'rgba(19,64,117,0.8)', 'rgba(19,64,117,1)'],
        credits: {
            enabled: false
        },
        mapNavigation: {
            enabled: true
        },
        colorAxis: {
            min: 0,
            stops: [
                [0, '#EFEFFF'],
                [0.5, Highcharts.getOptions().colors[0]],
                [1, Highcharts.color(Highcharts.getOptions().colors[0]).brighten(-0.5).get()]
            ]
        },

        series: [{
            name: 'Basemap',
            mapData: mapDataWorld,
            // joinBy: ['hc-key', 'key'],
            showInLegend: false,
            dataLabels: {
                enabled: true,
                format: "{point.name}"
            }
        }]
    }

    return (
        <div className="chart-container">
            <div className="chart-item">
                <HighchartsReact highcharts={Highcharts}
                                 constructorType = { 'mapChart' }
                                 options={mapOptions}/>
            </div>
        </div>
    );
};

export default MapHCH;

