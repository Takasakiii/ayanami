function indexInitializer() {
    return {
        filename: null,
        password: false,
        response: null,
        progress: false,

        upload() {
            this.progress = 0

            const form = this.$refs.form
            const data = new FormData(form)
            const xhr = new XMLHttpRequest()

            xhr.upload.addEventListener('loadstart', () => {
                this.progress = 0
            })

            xhr.upload.addEventListener('progress', (e) => {
                if(!e.lengthComputable) return
                this.progress = Math.round((e.loaded / e.total) * 100)
            })

            xhr.upload.addEventListener('loadend', () => {
                this.progress = false
            })

            xhr.addEventListener('load', () => {
                this.response = JSON.parse(xhr.responseText)
                this.password = data.get('password') ?? false
            })

            xhr.upload.addEventListener('error', () => {
                alert('Falha ao enviar o arquivo')
                this.progress = false
            })

            xhr.open(form.method, form.action)
            xhr.send(data)
        }
    }
}