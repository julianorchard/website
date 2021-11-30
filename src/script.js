// Create Leading Characters
	window.onload = function() {
		var leadingCharacters = document.getElementsByClassName("leading-character");
		for (var i = 0; i < leadingCharacters.length; i++) {
			var leadingCharacter        = leadingCharacters[i];
			var leadingCharacterContent = leadingCharacter.innerHTML;
			var node     = document.createElement("span");
			var textnode =  document.createTextNode(leadingCharacterContent.charAt(0));
			node.appendChild(textnode);
			leadingCharacter.appendChild(node);
		}
	}
