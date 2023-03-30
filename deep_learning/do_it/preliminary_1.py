from mxnet import np, npx

npx.set_np()

x = np.arange(12)

print(x)
print(x.shape)
print(x.size)

X = x.reshape(3, 4)
print(X)
