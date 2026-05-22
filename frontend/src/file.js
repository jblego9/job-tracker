export async function createFileHandleCSV() {
    const options = {
        types: [{
            description: "CSV Files",
            accept: { "text/csv": [".csv"] },
    }]};

    const handle = await window.showSaveFilePicker(options);
    return handle;
}

export async function writeToFile(fileHandle, contents) {
    const writable = await fileHandle.createWritable();
    await writable.write(contents);
    await writable.close();
}
