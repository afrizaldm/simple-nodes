<template>
    <ul class="tree">
        <li v-for="(node, index) in data" :key="index" class="tree-item">
            <div class="tree-item-content" @click="toggle(node)">
                <span class="toggle-icon">
                    {{ node.isToggle ? '+' : '-' }}
                </span>
                <span class="tree-item-name">{{ node?.name }}</span>
            </div>
            <div class="tree-item-actions">
                <button @click="edit(node)" class="tree-item-action">Edit</button>
                <button @click="remove(node)" class="tree-item-action">Hapus</button>
                <button @click="add(node)" class="tree-item-action">Tambah</button>
            </div>
            <Tree v-if="!node.isToggle" :data="(node.children as any)"></Tree>
        </li>
    </ul>
</template>
<script setup lang="ts">
import { Node } from "../types/Node";
import { axios } from "../modules/axios";
import { load } from "../services/node";
import router from "../router";

interface Props {
    data: Node[]
}

let { data } = defineProps<Props>()

const toggle = (item: Node) => {
    item.isToggle = !item.isToggle
}

const edit = (item: Node) => {
    
    router.push("edit/" + item.id)
}

const remove = async (item: Node) => {
    const res = await axios.delete('nodes/' + item.id)

    console.log(res)

    load()
}

const add = (item: Node) => {
    router.push("add/" + item.id)

}

const emit = defineEmits<{
    (e: 'add', node: Node): void;
    (e: 'remove', node: Node): void;
    (e: 'edit', node: Node): void;
}>();


</script>
<style scoped lang="scss">
.tree {
    list-style: none;
    padding: 0;
    margin: 0;
}

.tree-item {
    position: relative;
    margin-left: 20px;
    padding-left: 10px;
    border-left: 1px dashed #ddd;

    &:hover .tree-item-actions {
        opacity: 1;
        visibility: visible;
    }
}

.tree-item-content {
    display: flex;
    align-items: center;
    gap: 10px;
    cursor: pointer;
    transition: background-color 0.3s ease;

    &:hover {
        background-color: #f5f5f5;
    }
}

.toggle-icon {
    font-weight: bold;
    width: 20px;
    text-align: center;
    color: #666;
}

.tree-item-name {
    font-weight: 500;
    color: #333;
}

.tree-item-actions {
    margin-top: 5px;
    display: flex;
    gap: 5px;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.3s ease, visibility 0.3s ease;

    .tree-item-action {
        padding: 3px 8px;
        font-size: 12px;
        color: #fff;
        background-color: #007bff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        transition: background-color 0.3s ease;

        &:hover {
            background-color: #0056b3;
        }

        &:first-of-type {
            background-color: #28a745;

            /* Green for Edit */
            &:hover {
                background-color: #218838;
            }
        }

        &:nth-of-type(2) {
            background-color: #dc3545;

            /* Red for Remove */
            &:hover {
                background-color: #c82333;
            }
        }

        &:last-of-type {
            background-color: #17a2b8;

            /* Teal for Add */
            &:hover {
                background-color: #138496;
            }
        }
    }
}
</style>