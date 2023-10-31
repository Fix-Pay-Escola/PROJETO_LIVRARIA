function openModal(){
    const clicked = document.getElementById('janela_modal')
    clicked.classList.add('open')
    clicked.addEventListener('click', (e) =>{
        if(e.target.id == 'close' || e.target.id == 'janela_modal'){
            clicked.classList.remove('open')
        }
    })
}