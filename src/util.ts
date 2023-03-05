import { Options, IllustInfo } from "./v13s.types"

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
    history.pushState("", "", "/")

    const illustInfoEditDialog = document.getElementById("illustInfoEditDialog") as HTMLDivElement
    const bodyFilter = document.getElementById("bodyFilter") as HTMLDivElement
    illustInfoEditDialog.style.animationName = "disableModal"
    bodyFilter.style.animationName = "disableBodyFilter"
}