import { Options } from "./v13s.types"

export async function getOptions(): Promise<Options> {
    const response = await fetch("/api/options/")
    const options = (await response.json()) as Options
    return options
}
