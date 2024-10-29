export type Node = {
    id: any
    name: string
    parent_id: any
    children?: Node[],
}

export default Node