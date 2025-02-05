document.querySelectorAll('.nav-item').forEach((item) => {
    item.addEventListener('click', function (event) {
        // Remove the 'selected' class from all items
        document.querySelectorAll('.nav-item').forEach((i) => {
            i.classList.add('text-gray-400');
            i.classList.remove(
                'text-slate-200',
                'font-bold',
                'flex',
                'items-center'
            );
            const arrow = i.querySelector('span');
            if (arrow) arrow.remove();
        });

        let targetElement = event.target;
        console.log(targetElement);
        targetElement.classList.remove('text-gray-400');
        targetElement.classList.add(
            'text-slate-200',
            'font-bold',
            'flex',
            'items-center',
            'nav-item',
            'cursor-pointer'
        );

        // Add the 'selected' class to the clicked item and add arrow
        const arrowSpan = document.createElement('span');
        arrowSpan.textContent = '→';
        arrowSpan.classList.add('mr-2');
        let anchorTags = this.getElementsByTagName('a');
        anchorTags[0].prepend(arrowSpan);
    });
});
