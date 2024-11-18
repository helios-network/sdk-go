VERSION=${VERSION:-"v0.50.10-helios-8"}

# version must be start by v8

echo "Deploy sdk-go"
git add .
git commit -m "Publish $VERSION"
git push
git tag $VERSION
git push origin $VERSION
sleep 5
GOPROXY=proxy.golang.org go list -m github.com/Helios-Chain-Labs/sdk-go@$VERSION

echo "Publish done"