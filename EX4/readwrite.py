import addressbook_pb2

p1 = addressbook_pb2.Person()

p1.name = "Steve"
p1.id = 123
p1.email = "steve@yahoo.com"

phone_number1 = p1.phones.add()
phone_number1.number = "12345678"
phone_number1.type = addressbook_pb2.Person.MOBILE

phone_number2 = p1.phones.add()
phone_number2.number = "98765432"
phone_number2.type = addressbook_pb2.Person.HOME

with open("simple.bin", "wb") as f:
    bytesToString = p1.SerializeToString()
    f.write(bytesToString)

with open("simple.bin", "rb") as f:
    read_msg = addressbook_pb2.Person.FromString(f.read())
    print(read_msg)