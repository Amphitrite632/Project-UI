export type Options = {
    siteName: string?
}

export type RawIllust = {
    hash: string
    originalURL: string
    tags: string[]
}

export type RawTag = {
    type: "WORK" | "CHARACTER" | "AUTHOR"
    officialName: string
    aliases: string[]
    children?: string[]
}

export type IllustInfo = {
    illustID: string
    originalURL: string | null
    heightRatio: number
    tags: Tag[]
}

export type TagInfo = {
    tagID: string
    type: "WORK" | "CHARACTER" | "AUTHOR"
    name: string
}
