#include <iostram>

class Cat {
    public: 
        string name;
        int paw = 4;
    private:
        void Cat(string Name,int Paw){
            name = Name;
            paw=Paw;
        }   
}

void main(){
    Cat ronnie = new Cat("Ronnie",4);
   

}