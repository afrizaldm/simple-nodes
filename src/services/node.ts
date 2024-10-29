import { ref } from "vue";

import { Node } from "../types/Node";

import { axios } from "../modules/axios";

export const nodes = ref<Node>()

export const buildTree = (data: Node[]): Node[] => {
    const map: { [key: number]: Node } = {};
    const roots: Node[] = [];

    data.forEach(item => {
        map[item.id] = { ...item, children: [] }; 
    });

    data.forEach(item => {
        if (item.parent_id === 0) {
            roots.push(map[item.id]);
        } else {
            const parent = map[item.parent_id];
            if (parent) {
                parent.children!.push(map[item.id]);
            }
        }
    });

    return roots;
}

export const load = async () => {

    const res = await axios.get('/nodes')

    if (res?.data && res?.data.status == 'success') {

        const _nodes = res.data.data

        const tree: Node[] = buildTree(_nodes)

        console.log(tree)
        console.log(res?.data.data)

        nodes.value = tree as any
    }
}