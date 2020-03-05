old=v2.2.0
new=v2.2
git tag $new $old
git tag -d $old
git log --pretty=oneline