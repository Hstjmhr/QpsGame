import {defineComponent, h, createApp, ref} from "vue";
import type {Component} from "vue";
import Gvb_message_record_modal from "@/components/common/gvb_message_record_modal.vue";


const visible = ref<undefined | boolean>(undefined)
const userIDRef = ref<number>(0)
const nickNameRef = ref<string>("")
const messageRecordRef = ref()

export function showMessageRecord(userID: number, nickName: string) {

    if (visible.value === undefined) {
        const component = defineComponent({
            setup() {
                const onClose = () => {
                    visible.value = false
                }


                return {
                    visible,
                    userIDRef,
                    nickNameRef,
                    onClose
                }
            },
            render() {
                return h(Gvb_message_record_modal as Component, {
                    visible: this.visible,
                    userID: this.userIDRef,
                    nickName: this.nickNameRef,
                    onClose: this.onClose,
                    ref: messageRecordRef,
                })
            }
        })

        const app = createApp(component)
        const div = document.createElement("div")
        app.mount(div)
        document.body.appendChild(div)
    }

    if (userID === userIDRef.value) {
        messageRecordRef.value.getData()
    }
    visible.value = true
    userIDRef.value = userID
    nickNameRef.value = nickName
}