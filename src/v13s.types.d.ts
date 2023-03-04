export type Options = {
    siteName: string
    siteDescription: string
}

export type IllustInfo = {
    illustID: string
    hash: string
    originalURL: string
    heightRatio: string
    tags: TagInfo[]
    createdAt: string
}

export type TagInfo = {
    tagID: string
    type: "WORK" | "CHARACTER" | "AUTHOR"
    officialName: string
}
