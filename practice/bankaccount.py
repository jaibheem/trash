class BankAccount(object):
	def __init__(self, balance=0):
		self.balance = 0
	def withdraw(self, amount):
		self.balance -= amount
		return self.balance
	def deposit(self, amount):
		self.balance += amount
		return self.balance


#print a.deposit(20)
#print b.deposit(30)

#print a.withdraw(10)
#print b.withdraw(5)

#print "Balance in A account is: ", a.balance
#print "Balance in B account is: ", b.balance

class MinimumBalanceAccount(BankAccount):
	def __init__(self, min_balance):
		BankAccount.__init__(self)
		self.min_balance = min_balance

	def withdraw(self, amount):
		if self.balance - amount < self.min_balance:
			print "Sorry, minimum balance must be maintained"
		else:
			return BankAccount.withdraw(self, amount)


#Creating a instance to have a minimum balance of 10 rs
a = MinimumBalanceAccount(10)
#Depositing 30 rs to account
a.deposit(30)
#b = BankAccount()

#with drawing 5 rs
print a.withdraw(5)

print a.balance