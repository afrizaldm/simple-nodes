<template>
    <div>
        <button @click="back">kembali</button>
        <h2>Add Node</h2>
        <form @submit.prevent="addNode">
            <div>
                <label for="name">Node Name:</label>
                <input type="text" v-model="name" required />
            </div>
            <button type="submit">Simpan</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '../router';
import { axios } from '../modules/axios';
import { useRoute } from 'vue-router';

const route = useRoute()

const name = ref('');

const emit = defineEmits<{
    (e: 'add-node', node: { name: string; parent_id: number }): void;
}>();

const addNode = () => {

    const parentId: any = Number(route.params?.id) ?? 0

    axios.post('nodes', {
        name: name.value,
        parent_id: parentId,
    })


};

const back = () => {
    router.back()
}
</script>

<style scoped>
form {
    margin-bottom: 20px;
}
</style>