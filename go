#!/bin/bash
touch $1
echo '#!/bin/bash' > $1
echo $2 >> $1

chmod u+x $1

git add $1
git commit -m "solve task ${$1:0:1}"
git push
