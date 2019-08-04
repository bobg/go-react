bundle.js: component.js
	-rm -f $@
	TEMP=`mktemp --suffix .js tempXXXX` && \
	npx babel $< > $$TEMP && \
	npx browserify $$TEMP > $@ && \
	rm -f $$TEMP
