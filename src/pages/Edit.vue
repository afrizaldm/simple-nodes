<template>
    <div>
        <button @click="back">kembali</button>
        <h2>Add Node</h2>
        <form @submit.prevent="editNode">
            <div>
                <label for="name">Node Name:</label>
                <input type="text" v-model="name" required />
            </div>
            <button type="submit">Simpan</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import router from '../router';
import { axios } from '../modules/axios';
import { useRoute } from 'vue-router';
import { load } from "../services/node";

const route = useRoute()

const name = ref('');
const parent_id = ref<number>();

const emit = defineEmits<{
    (e: 'add-node', node: { name: string; parent_id: number }): void;
}>();

const editNode = async () => {
    
    const id: any = Number(route.params?.id) ?? 0

    await axios.put('nodes/' + id, {
        name: name.value,
        parent_id: parent_id.value,
    })

    router.back()
    
    load()

};

const back = () => {
    router.back()
}

onMounted(async() => {

    const id: any = Number(route.params?.id) ?? 0

    const res = await axios.get('nodes/' + id)

    console.log(res)

    name.value = res.data.data.name ?? ''
    parent_id.value = res.data.data.parent_id ?? 0

})
</script>

<style scoped>
form {
    margin-bottom: 20px;
}
</style>