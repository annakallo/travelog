import http from './httpService';

const  apiEndpoint = '/api/countries';


function countryUrl(id) {
    return `${apiEndpoint}/${id}`;
}

export function getCountries() {
    return http.get(apiEndpoint);
}

export function deleteCountry(id) {
    return http.delete(countryUrl(id));
}

export function saveEntry(country) {
    return http.post(apiEndpoint, country);
}
