// :arg backend: "scheme://host:port",
// :arg filter: map with keys:
async function FetchBackendRows(backend, path, filter) {
	let rets = [];
	let query = "";
	if (filter) {
		let queries = [];
		for (let key in filter) {
			if (!filter.hasOwnProperty(key)) {
				continue
			}
			queries.push(`${key}=${filter[key]}`)
		}
		query = "?" + queries.join("&")
	}
	let url = `${backend}${path}${query}`;
	console.log(`fetching ${url}`);
	await fetch(url, {method: "GET", cache: 'no-cache',}).then(
		async (response) => {
			if (!response.ok) {
				throw new Error(response.statusText)
			}
			let respObj = await response.json();
			rets.push(respObj)
		}).catch(err => {
		console.log(`error fetch ${url}: ${err}`);
		rets.push({Error: `error fetch ${url}: ${err}`})
	});
	return rets[0];
}

function GetBackend() {
	let backend = window.location.origin; // GUI and API have a same scheme://host:port
	if (backend === "" || backend === "null") {
		backend = "http://127.0.0.1:39539";
	}
	return backend
}
