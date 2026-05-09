const backendUrl = "";
export const apiUrl = `${backendUrl}/api/v1/applications`;

export async function getDataJson(url) {
    const response = await fetch(url, { method: "GET" });

    if (!response.ok)
        throw new Error(`Failed to get JSON data from '${url}'. Response status: ${response.status}`);

    return await response.json();
}

export async function putData(url, data) {
    const response = await fetch(url, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    });

    if (!response.ok)
        throw new Error(`Failed to put data to '${url}'. Response status: ${response.status}. Data: ${data}`);
}

export async function postData(url, data) {
    const response = await fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    });

    if (!response.ok)
        throw new Error(`Failed to post data do '${url}'. Response status: ${response.status}. Data: ${data}`);
}

export async function deleteData(url) {
    const response = await fetch(url, { method: "DELETE" });

    if (!response.ok)
        throw new Error(`Failed to delete data at '${url}'. Response status: ${response.status}`);
}
