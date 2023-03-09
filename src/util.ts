import { Options } from "./v13s.types"

export async function getOptions(): Promise<Options> {
    const response = await fetch("/api/options/")
    const options = (await response.json()) as Options
    return options
}

export function editIllustInfo() {
    history.pushState("", "", "/edit/illust/")
    activateIllustInfoEditDialog()
}

export function activateIllustInfoEditDialog() {
    const illustInfoEditDialog = document.getElementById("illustInfoEditDialog") as HTMLDialogElement
    const bodyFilter = document.getElementById("bodyFilter") as HTMLDivElement

    illustInfoEditDialog.open = true
    bodyFilter.style.animationName = "activateBodyFilter"
    illustInfoEditDialog.style.animationName = "activateModal"
}

export function disableIllustInfoEditDialog() {
    history.replaceState("", "", "/")

    const illustInfoEditDialog = document.getElementById("illustInfoEditDialog") as HTMLDialogElement
    const bodyFilter = document.getElementById("bodyFilter") as HTMLDivElement
    const searchBox = document.getElementById("searchBox") as HTMLInputElement
    illustInfoEditDialog.style.animationName = "disableModal"
    if (document.activeElement != searchBox) {
        bodyFilter.style.animationName = "disableBodyFilter"
    }
    setTimeout(function () {
        illustInfoEditDialog.open = false
    }, 375)
}
