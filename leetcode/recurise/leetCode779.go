package main


func kthGrammar(N int, K int) int {
   if N == 1 && K == 1 {
   	 return 0
   }
   result := 0
   if K % 2 ==0 {
   	  result = kthGrammar(N-1, K /2)
   	  if result == 0 {
   	  	result = 1
	  } else {
	  	result = 0
	  }
   } else {
   	  result = kthGrammar(N-1, K /2 +1)
   	  if result == 0 {
   	  	result = 0
	  } else  {
	  	result = 1
	  }
   }
   return result
}