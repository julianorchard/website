// Create Leading Chars
	window.onload = function() {
		var larchars = document.getElementsByClassName("larchar");
		for (var i = 0; i < larchars.length; i++) {
			var larchar = larchars[i];
			var larcharContent = larchar.innerHTML;

			var node = document.createElement("span");
			var textnode =  document.createTextNode(larcharContent.charAt(0));
			node.appendChild(textnode);
			larchar.appendChild(node);
		}
	}
// Scrolling Menu
	window.onscroll = function() {
		scrollMenu2()
	};
	function scrollMenu2() {
		var navBar = document.getElementsByTagName("ul")[0];
		var hero = document.getElementById("hero");
		if ( document.body.scrollTop > 450 || document.documentElement.scrollTop > 450 ) {
			navBar.className = "fixed";
			hero.style.marginTop = "4.5%";
			// id hero padding top 10%;
		} else {
			navBar.className = "";
			hero.style.marginTop = "0%";
		}
	}

