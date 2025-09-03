import results from './results.json'

function createRow(container, results) {
    const rowHeader = document.createElement("div");
    rowHeader.textContent = results[0].condition_1;
    rowHeader.className = 'row-header';
    container.appendChild(rowHeader);

    results.forEach(result => {
        const button = document.createElement("button");
        button.className = "outline contrast";
        button.textContent = `${result.words.length} words`;

        button.addEventListener("click", () => {
            const modal = document.getElementById("modal");
            modal.setAttribute("open", "");

            const modalTitle = document.getElementById("modal-title");
            modalTitle.textContent = `"${result.condition_1}" & "${result.condition_2}"`;

            const modalList = document.getElementById("modal-list");
            modalList.innerHTML = "";
            result.words.forEach(word => {
                const listItem = document.createElement("li");
                listItem.textContent = word;
                modalList.appendChild(listItem);
            });
        });

        container.appendChild(button);
    });
}

function createColumnHeaders(container, results) {
    const blank = document.createElement("div");
    container.appendChild(blank);

    results.forEach(result => {
        const header = document.createElement("div");
        header.className = 'col-header';
        header.textContent = result.condition_2;
        container.appendChild(header);
    });
}

let container = document.getElementById("results");
let gameInfo = document.getElementById("game-info");

gameInfo.textContent = `Game #${results.game_number} - ${new Date(results.timestamp).toLocaleDateString()}`;

createColumnHeaders(container, results.results.slice(0, 3));
createRow(container, results.results.slice(0, 3));
createRow(container, results.results.slice(3, 6));
createRow(container, results.results.slice(6, 9));

// Sort alphabetically
const buttonSort1 = document.getElementById("modal-sort-1");
buttonSort1.addEventListener("click", () => {
    const modalList = document.getElementById("modal-list");
    const items = Array.from(modalList.children);
    items.sort((a, b) => a.textContent.localeCompare(b.textContent));
    modalList.innerHTML = "";
    items.forEach(item => modalList.appendChild(item));
});

// Sort by length
const buttonSort2 = document.getElementById("modal-sort-2");
buttonSort2.addEventListener("click", () => {
    const modalList = document.getElementById("modal-list");
    const items = Array.from(modalList.children);
    items.sort((a, b) => b.textContent.length - a.textContent.length);
    modalList.innerHTML = "";
    items.forEach(item => modalList.appendChild(item));
});

const modalCloseBtn = document.getElementById("modal-close");
modalCloseBtn.addEventListener("click", () => {
    const modal = document.getElementById("modal");
    modal.removeAttribute("open");
});