from novaclient import client as nova
from novaclient.v2 import client

if __name__ == "__main__":
    # client : client.Client = client.Client()
    # client.server_snapshots.
    nova = nova.Client('2.1')
    print(type(nova))
    ss = dir(nova)

    # 区分不同类型的属性
    methods = {
        'instance_methods': [],
        'class_methods': [],
        'static_methods': [],
    }
    fields = []
    for attr in ss:
        if not attr.startswith('_'):
            value = getattr(nova, attr)
            if isinstance(value, client.Client):
                methods['instance_methods'].append(attr)
            elif isinstance(value, classmethod):
                methods['class_methods'].append(attr)
            elif isinstance(value, staticmethod):
                methods['static_methods'].append(attr)
            else:
                fields.append(attr)


    for k,v in methods.items():
        print(k)
        print(v)
        print("*"*10)