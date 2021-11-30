// Create Leading Characters
	window.onload = function() {
		// If you want this for every paragraph (which I think
		// looks a little too much), put it in a loop to the length
		// of the leadingCharacters const.
		const leadingCharacters = document.getElementsByTagName("p");
		let leadingCharacter        = leadingCharacters[0];
		let leadingCharacterContent = leadingCharacter.innerHTML.toLowerCase();
		let node     = document.createElement("span");
		let textnode =  document.createTextNode(leadingCharacterContent.charAt(0));
		// To Lower
		node.appendChild(textnode);
		leadingCharacter.appendChild(node);
	}
